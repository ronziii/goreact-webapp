import axios from 'axios'

const BASE_URL = 'http://localhost:3001/api/v1/'

export const GET = (endpoint, params) => {
  return axios.get(BASE_URL + endpoint, {params})
}

export const POST = (endpoint, payload) => {
return axios.post(BASE_URL + endpoint, payload)
}

export const DELETE = (endpoint, params) => {
  return axios.delete(BASE_URL + endpoint, {params})
}