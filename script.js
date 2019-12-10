const express = require("express");
const Joi = require("joi");
const app = express();
// fetch the local JSON files
const bands = require("./data/bands.json");
const locations = require("./data/location.json");

// use json files
app.use(express.json());

// C -> Creat -> POST
// R -> Read -> GET
// U -> Update -> PUT
// D -> Delete -> Delet

// read request handlers
// request is what is send from the client side
// respose is what is send from the server side
// display the message when the URL consist of '/'
app.get("/", (req, res) => {
  res.send("Welcome to band groupi REST API");
});

// display the list of bands when URL consists of api bands
app.get("/api", (req, res) => {
  const api = [
    {bands:"http://localhost:8080/api/bands"},
    {locations:"http://localhost:8080/api/locations"}
  ]
  res.send(api)
});

// display the list of bands when URL consists of api locations
app.get("/api/locations", (req, res) => {
  res.send(locations);
});

// display the list of bands when URL consists of api bands
app.get("/api/bands", (req, res) => {
  res.send(bands);
});

// display the information of specific ______ when you mention the id
app.get("/api/location/:id", (req, res) => {
  const location = locations.index.find(b => b.id === parseInt(req.params.id));
  // if there is no valid band ID, then diplay an error with the folowing
  error404(location)
  res.send(location);
});

// display the information of specific ______ when you mention the id
app.get("/api/bands/:id", (req, res) => {
  const band = bands.index.find(b => b.id === parseInt(req.params.id));
  // if there is no valid band ID, then diplay an error with the folowing
 error404(band)
  res.send(band);
});

// CREATE request handler
// create new band information
app.post("/api/bands", (req, res) => {
  const { error } = validateCustomer(req.body);
  if (error) {
    res.status(400).send(error.details[0].message);
    return;
  }
  // add band
  const band = {
    if: bands.length + 1,
    title: req.body.title
  };
  bands.push(band);
  res.send(band);
});

// PUT method update the resurce
app.put("/api/bands/:id", (req, res) => {
  const band = bands.find(c => c.id === parseInt(req.params.id));
  error404(band)
  // the band is found so...
  const { error } = validateCustomer(req.body);
  if (error) {
    res.status(400).send(error.details[0].message);
    return;
  }

  band.title = req.body.title;
  res.send(band);
});

// delete request handler
// delete band details
app.delete("/api/bands/:id", (req, res) => {
  const band = bands.find(c => c.id === parseInt(req.params.id));
  error404(band)

  const index = bands.indexOf(band);

  // the element is deleted and the slice is push ->>
  bands.splice(index, 1);

  res.send(bands);
});

const error404 = value => {
  if (!value)
    res
      .status(404)
      .send(
        '<h2 style="font-family: Malgun Gothic; color: darkred;"> Ooops... error on query</h2>'
      );
}

// validates the information to add to the JSON
const validateCustomer2 = band => {
  const schema = {
    title: Joi.string()
      .min(3)
      .required()
  };
  return Joi.validate(band, schema);
};

// assign a port for the server 8080
const port = process.env.PORT || 8080;
app.listen(port, () => console.log(`listening on port ${port}...`));
