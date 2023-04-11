const express = require("express");
require("dotenv").config();
const app = express();
app.use(express.json());

app.get("/", (req, res) => {
  res.send("Hello, World!");
});

app.listen(process.env.PORT, (err) => {
  if (err) throw err;
  console.log(`Server is running on port ${process.env.PORT}`);
});
