import { Fragment, useState, useEffect } from "react";
import RecipeList from "../../components/RecipesList/RecipeList";
import SearchBar from "../../components/SearchBar/SearchBar";
import axios from "axios";

const HomePage = () => {
  const [recipes, setRecipes] = useState([]);

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
        console.log(response.data.recipes);
        setRecipes(response.data.recipes);
      } catch (error) {
        console.error(error.response.data.error);
      }
    };
    fetchRecipes();
  }, []);

  return (
    <Fragment>
      <SearchBar />
      <RecipeList recipes={recipes} showButton={false} />
    </Fragment>
  );
};

export default HomePage;
