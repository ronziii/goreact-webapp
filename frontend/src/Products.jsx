import React from "react";
import TableContainer from "@material-ui/core/TableContainer";
import {Paper} from "@material-ui/core";
import Table from "@material-ui/core/Table";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import TableBody from "@material-ui/core/TableBody";
import {DeleteOutline} from '@material-ui/icons';
import _ from 'lodash'
import {DELETE, GET, POST} from "./http";
import './products.css'
import IconButton from "@material-ui/core/IconButton";
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogContent from "@material-ui/core/DialogContent";
import TextField from "@material-ui/core/TextField";
import DialogActions from "@material-ui/core/DialogActions";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert from '@material-ui/lab/Alert';

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

export default class Products extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      open: false,
      snackBarOpen: false,
      snackBarTypeError: true,
      name: '',
      price: null,
      message: '',
      products: []
    };
  }

  deleteProduct = (productId) => {
    DELETE(`products/${productId}`, {})
      .then(resp => {
        let products = _.filter(this.state.products, function (product) {
          return product.id !== productId;
        });
        this.setState({...this.state, products})
        this.handleOpenSnackBar("Product successfully deleted", false)
      })
      .catch(error => {
        this.handleOpenSnackBar("Failed to delete product", true)
      })
  }

  createProduct = () => {
    POST('products', {name: this.state.name, price: this.state.price})
      .then(resp => {
        this.setState({...this.state, name: '', price: null, products: [...this.state.products, resp.data]})
        this.handleClose()
        this.handleOpenSnackBar("Product successfully created", false)
      })
      .catch(error => {
        this.handleOpenSnackBar("Failed to create product", true)
        this.handleClose()
      })
  }

  handleClickOpen = () => {
    this.setState({...this.state, open: true})
  }

  handleClose = () => {
    this.setState({...this.state, open: false})
  }

  handleOpenSnackBar = (message, error) => {
    this.setState({...this.state, message, snackBarTypeError: error, snackBarOpen: true})
  }

  handleCloseSnackBar = () => {
    this.setState({...this.state, snackBarOpen: false, message: ''})
  }

  componentDidMount() {
    GET("products", {}).then(
      resp => {
        this.setState({...this.state, products: resp.data.products})
      }
    ).catch(error => {
      this.handleOpenSnackBar("Failed to fetch products", true)
    })
  }

  render() {
    return (
      <div className="productsContainer">
        <Snackbar anchorOrigin={{vertical: 'top', horizontal: 'center'}} open={this.state.snackBarOpen}
                  autoHideDuration={4000} onClose={this.handleCloseSnackBar}>
          <Alert onClose={this.handleClose} severity={this.state.snackBarTypeError ? 'error' : 'success'}>
            {this.state.message}
          </Alert>
        </Snackbar>

        <Button variant="contained" size="medium" color="primary" style={{float: 'right'}}
                onClick={() => this.handleClickOpen()}>
          Create Product
        </Button>
        <TableContainer component={Paper} className="tableContainer">
          <Table aria-label="customized table" className="table">
            <TableHead>
              <TableRow>
                <TableCell>Id</TableCell>
                <TableCell align="center">Name</TableCell>
                <TableCell align="center">Price </TableCell>
                <TableCell align="center"> </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {this.state.products.map((product) => (
                <TableRow key={product.id}>
                  <TableCell component="th" scope="row">
                    {product.id}
                  </TableCell>
                  <TableCell align="center">{product.name}</TableCell>
                  <TableCell align="center">{product.price}</TableCell>
                  <TableCell align="center">
                    <IconButton onClick={() => this.deleteProduct(product.id)} color="secondary" aria-label="delete">
                      <DeleteOutline/>
                    </IconButton>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>

        <Dialog open={this.state.open} onClose={this.handleClose} aria-labelledby="form-dialog-title">
          <DialogTitle id="form-dialog-title">Create Product</DialogTitle>
          <DialogContent>
            <TextField
              autoFocus
              margin="dense"
              id="name"
              label="Product Name"
              type="text"
              value={this.state.name}
              onChange={e => this.setState({...this.state, name: e.target.value})}
              fullWidth
            />
            <TextField
              margin="dense"
              id="price"
              label="Product Price"
              type="number"
              value={this.state.price}
              onChange={e => this.setState({ ...this.state, price: parseFloat(e.target.value) })}
              fullWidth
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose} color="primary">
              Cancel
            </Button>
            <Button onClick={this.createProduct} color="primary" disabled={this.state.name.length < 2}>
              Create
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}