const express = require("express");
const Joi = require("joi");
const app = express();

// fetch the local JSON files
const artists = require("./../data/artists.json");
const locations = require("./../data/location.json");
const dates = require("./../data/dates.json");

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
  res.send("<h1> Welcome to band groupi REST API </h1>");
});

// display the list of artists when URL consists of api artists
app.get("/api", (req, res) => {
  const api = [
    { artists: "http://localhost:8080/api/artists" },
    { locations: "http://localhost:8080/api/locations" },
    { dates: "http://localhost:8080/api/dates" }
  ];
  res.send(api);
});

// display the list of artists when URL consists of api locations
app.get("/api/locations", (req, res) => {
  res.send(locations);
});

app.get("/api/artists", (req, res) => {
  res.send(artists);
});

app.get("/api/dates", (req, res) => {
  res.send(dates);
});

// display the information of specific ______ when you mention the id
app.get("/api/locations/:id", (req, res) => {
  const location = locations.index.find(b => b.id === parseInt(req.params.id));
  // if there is no valid band ID, then diplay an error with the folowing
  error404(location);
  res.send(location);
});

app.get("/api/artists/:id", (req, res) => {
  const band = artists.index.find(b => b.id === parseInt(req.params.id));
  error404(band);
  res.send(band);
});

app.get("/api/dates/:id", (req, res) => {
  const date = date.index.find(d => d.id === parseInt(req.params.id));
  error404(date);
  res.send(date);
});

// relation between the locations and the dates
app.get("/api/relation/", (req, res) => {
  res.send(date);
});

const getRelation = () => {
  locations.index.forEach(ele => {
    dates.index.forEach(dEle => {
      arrDates = locationNdate(dEle)
      if (ele.id === dEle.id) { 
        ele.location.forEach(lEle => {
          var index = [ { location: lEle, dates: d} ]
        })
       }
    });
  });
};
const locationNdate = (date) => {
  var d
  date.forEach(ele => {
    if (ele[0] == '*') {
      d.push(ele.replace('*',''))
    }
  })
}
// CREATE request handler
// create new band information
app.post("/api/artists", (req, res) => {
  const { error } = validateCustomer(req.body);
  if (error) {
    res.status(400).send(error.details[0].message);
    return;
  }
  // add band
  const band = {
    if: artists.length + 1,
    title: req.body.title
  };
  artists.push(band);
  res.send(band);
});

// PUT method update the resurce
app.put("/api/artists/:id", (req, res) => {
  const band = artists.find(c => c.id === parseInt(req.params.id));
  error404(band);
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
app.delete("/api/artists/:id", (req, res) => {
  const band = artists.find(c => c.id === parseInt(req.params.id));
  error404(band);

  const index = artists.indexOf(band);

  // the element is deleted and the slice is push ->>
  artists.splice(index, 1);

  res.send(artists);
});

const error404 = value => {
  if (!value)
    res
      .status(404)
      .send(
        '<h2 style="font-family: Malgun Gothic; color: darkred;"> Ooops... error on query</h2>'
      );
};

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
