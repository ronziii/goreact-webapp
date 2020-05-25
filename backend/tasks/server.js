/* global process */
var child_process = require('child_process');

module.exports = {
  start: function startServer(options) {
    options = options || {};
    var stdio = options.stdio || 'inherit';
    if (options.silent) {
      stdio = ['ignore','ignore','ignore'];
    }

    return child_process.spawn(process.env.APP_NAME, {
      stdio: stdio
    });
  }
};
