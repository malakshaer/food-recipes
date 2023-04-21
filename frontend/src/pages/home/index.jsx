import { Fragment, useState, useEffect } from "react";
import RecipeList from "../../components/RecipesList/RecipeList";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";
import axios from "axios";
import classes from "./Home.module.css";

const HomePage = () => {
  const [recipes, setRecipes] = useState([]);
  const [search, setSearch] = useState([]);
  const [result, setResult] = useState(0);
  const [query, setQuery] = useState("");

  useEffect(() => {
    const fetchRecipes = async () => {
      try {
        const token = localStorage.getItem("token");
        const headers = {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        };
        const response = await axios.get(`${process.env.API_ENDPOINT}recipe`, {
          headers,
        });
        // console.log(response.data.recipes);
        setRecipes(response.data.recipes);
      } catch (error) {
        console.error(error.response.data.error);
      }
    };
    fetchRecipes();
  }, []);

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const token = localStorage.getItem("token");
      const headers = {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      };
      const response = await axios.get(
        `${process.env.API_ENDPOINT}search?q=${query}`,
        { headers }
      );
      console.log(response.data.results.length);
      setResult(response.data.results.length);
      setSearch(response.data.results);
    } catch (error) {
      console.error(error.response.data.error);
    }
  };

  return (
    <Fragment>
      <form onSubmit={handleSubmit} className={classes.form}>
        <input
          type="text"
          placeholder="Search for a recipe"
          className={classes.input}
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />
        <button type="submit" className={classes.btn}>
          <FontAwesomeIcon icon={faSearch} />
        </button>
      </form>
      {result > 0 ? (
        <RecipeList
          recipes={search}
          showButton={false}
          showActionButton={true}
        />
      ) : (
        <RecipeList
          recipes={recipes}
          showButton={false}
          showActionButton={true}
        />
      )}
    </Fragment>
  );
};

export default HomePage;
