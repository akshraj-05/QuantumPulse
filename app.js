const express = require('express');
const path = require('path');
const pingWebsite = require('./ping/ping'); // Assuming you have a ping module

const app = express();
app.use(express.urlencoded({ extended: true }));
app.use(express.json());

// Serve static files (CSS, JS, etc.)
app.use('/static', express.static(path.join(__dirname, 'static')));

// Define route for root path ("/")
app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'templates', 'form.html'));
});

// Handle form submission
app.post('/', async (req, res) => {
    const domain = req.body.domain;
    console.log(domain);

    try {
        let obj = await pingWebsite(domain);
        res.send(
            `The latency of domain: ${domain} is ${obj.latency.toString()} ms`  
        );
    } catch (error) {
     
        res.status(404).send({
            success: false,
            message: error.message
        });
    }
});

// Start the server
const PORT = 80;
app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}`);
});
