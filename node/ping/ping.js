const ping = require('ping');

function pingWebsite(domain) {
  return ping.promise.probe(domain)
    .then(function (res) {
      if (res.alive) {
        return { alive: true, latency: res.time };
      } else {
        return { alive: false };
      }
    })
    .catch(function (err) {
      throw err;
    });
}

module.exports = pingWebsite;
