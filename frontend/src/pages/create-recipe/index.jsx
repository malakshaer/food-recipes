import { useState } from "react";
import classes from "./CreateRecipe.module.css";
import axios from "axios";

const CreateRecipe = () => {
  const [name, setName] = useState("");
  const [image, setImage] = useState(null);
  const [instructions, setInstructions] = useState("");
  const [ingredients, setIngredients] = useState([]);
  const [category, setCategory] = useState("");
  const [totalTime, setTotalTime] = useState(0);
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  const submitHandler = async (event) => {
    event.preventDefault();
    const token = localStorage.getItem("token");

    let recipeData = {
      name: name,
      instructions: instructions,
      ingredients: ingredients.split(";").map((ingredient, index) => ({
        id: String(index + 1),
        text: ingredient.trim(),
      })),
      recipecategory: category,
      totaltime: totalTime,
    };
    if (image) {
      recipeData = { ...recipeData, recipeimage: image };
    }

    try {
      const response = await axios.post(
        `${process.env.API_ENDPOINT}recipe`,
        recipeData,
        {
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        }
      );
      console.log(response.data);
      setSuccessMessage("Recipe created successfully");
    } catch (error) {
      console.log(error.response.data);
      setErrorMessage(error.response.data.error || "An error occurred");
    }
  };

  return (
    <div className={classes.card}>
      <form className={classes.form} onSubmit={submitHandler}>
        <h2>Unleash your creativity and let your taste buds dance with joy</h2>
        <div>
          <div className={classes.control}>
            <label htmlFor="name">Recipe Name:</label>
            <input
              type="text"
              required
              id="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <div className={classes.control}>
            <label htmlFor="image">Recipe Image:</label>
            <input
              type="file"
              id="image"
              name="image"
              onChange={(e) => {
                const file = e.target.files[0];
                const reader = new FileReader();
                reader.readAsDataURL(file);
                reader.addEventListener("load", () => {
                  const base64String = reader.result.split(",")[1];
                  setImage(base64String);
                });
              }}
            />
          </div>
          <div className={classes.control}>
            <label htmlFor="category">Category:</label>
            <select
              className={classes.category}
              id="category"
              required
              value={category}
              onChange={(e) => setCategory(e.target.value)}
            >
              <option value="">Select Category</option>
              <option value="Salads">Salads</option>
              <option value="Soups">Soups</option>
              <option value="Desserts">Desserts</option>
              <option value="Main Dishes">Main Dishes</option>
              <option value="Sides">Sides</option>
              <option value="Breakfast">Breakfast</option>
              <option value="Drinks">Drinks</option>
              <option value="Snacks">Snacks</option>
            </select>
          </div>
          <div className={classes.control}>
            <label htmlFor="totalTime">Total Time:</label>
            <input
              type="number"
              required
              id="totalTime"
              value={totalTime}
              onChange={(e) => setTotalTime(e.target.value)}
            />
          </div>
        </div>
        <div>
          <div className={classes.control}>
            <label htmlFor="ingredients">Ingredients:</label>
            <textarea
              id="ingredients"
              required
              rows="5"
              placeholder="Insert (;) after each ingredients"
              value={ingredients}
              onChange={(e) => setIngredients(e.target.value)}
            ></textarea>
          </div>
          <div className={classes.control}>
            <label htmlFor="instructions">Instructions:</label>
            <textarea
              id="instructions"
              required
              rows="5"
              value={instructions}
              onChange={(e) => setInstructions(e.target.value)}
            ></textarea>
          </div>
          <div className={classes.actions}>
            <button>Add Recipe</button>
          </div>
          {errorMessage && <div className={classes.error}>{errorMessage}</div>}
          {successMessage && (
            <div className={classes.success}>{successMessage}</div>
          )}
        </div>
      </form>
    </div>
  );
};

export default CreateRecipe;
