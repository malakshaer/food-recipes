import React, { useState } from "react";
import classes from "./ProfileNavBar.module.css";
import RecipeList from "../RecipesList/RecipeList";
import { useEffect } from "react";
import axios from "axios";

const ProfileNavBar = () => {
  const [listType, setListType] = useState("");
  const [isActive, setIsActive] = useState(false);
  const [recipes, setRecipes] = useState([]);
  const [savedRecipes, setSavedRecipes] = useState([]);

  useEffect(() => {
    const fetchRecipes = async () => {
      try {
        const token = localStorage.getItem("token");
        const headers = {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        };
        const response = await axios.get(`${process.env.API_ENDPOINT}user`, {
          headers,
        });

        const userRecipes = response.data.data.Recipes;
        setRecipes(userRecipes);
      } catch (error) {
        console.error(error);
      }
    };
    fetchRecipes();
  }, []);

  useEffect(() => {
    const fetchRecipes = async () => {
      try {
        const token = localStorage.getItem("token");
        const headers = {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        };
        const response = await axios.get(`${process.env.API_ENDPOINT}user`, {
          headers,
        });
        console.log(response.data.data.SavedRecipes);
        const userSavedRecipes = response.data.data.SavedRecipes;
        setSavedRecipes(userSavedRecipes);
      } catch (error) {
        console.error(error);
      }
    };
    fetchRecipes();
  }, []);

  const renderList = () => {
    switch (listType) {
      case "my-recipes":
        return (
          <RecipeList
            recipes={recipes}
            showButton={true}
            showActionButton={true}
          />
        );
      case "saved-recipes":
        return (
          <RecipeList
            recipes={savedRecipes}
            showButton={false}
            showUnActionButton={true}
          />
        );
      default:
        return (
          <RecipeList
            recipes={recipes}
            showButton={true}
            showActionButton={true}
          />
        );
    }
  };

  return (
    <div>
      <nav className={classes.users_navigation}>
        <div className={classes.users_navigation_menu}>
          <ul>
            <li>
              <button
                className={classes.isActive ? "active" : ""}
                onClick={() => {
                  setIsActive(true);
                  setListType("my-recipes");
                }}
              >
                My Recipes
              </button>
            </li>
            <li>
              <button
                className={classes.isActive ? "active" : ""}
                onClick={() => {
                  setIsActive(true);
                  setListType("saved-recipes");
                }}
              >
                Saved Recipes
              </button>
            </li>
          </ul>
        </div>
      </nav>
      {renderList()}
    </div>
  );
};

export default ProfileNavBar;
