const express = require('express');
const path = require('path');
const { ping } = require('./ping/ping'); // Assuming you have a ping module

const app = express();

// Serve static files (CSS, JS, etc.)
app.use('/static', express.static(path.join(__dirname, 'static')));

// Define route for root path ("/")
app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'templates', 'form.html'));
});

// Handle form submission
app.post('/', (req, res) => {
    const domain = req.body.domain;
    console.log(domain);

   // ping(domain)
   
});

// Start the server
const PORT = 80;
app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}`);
});
