import { useState } from "react";
import classes from "./EditRecipe.module.css";
import axios from "axios";
import { useRouter } from "next/router";

const EditRecipe = () => {
  const router = useRouter();
  const { id } = router.query;
  const [name, setName] = useState();
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
      ingredients:
        ingredients.length > 0
          ? ingredients.split(";").map((ingredient, index) => ({
              id: String(index + 1),
              text: ingredient.trim(),
            }))
          : [],
      recipecategory: category,
      totaltime: String(totalTime),
    };

    if (image) {
      recipeData = { ...recipeData, recipeimage: image };
    }

    try {
      const response = await axios.put(
        `${process.env.API_ENDPOINT}recipe/${id}`,
        recipeData,
        {
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        }
      );
      console.log(response);
      setSuccessMessage("Recipe updated successfully");
      setErrorMessage(null);
    } catch (error) {
      console.log(error);
      setErrorMessage(error.response.data.error || "An error occurred");
      setSuccessMessage(null);
    }
  };
  const deleteRecipe = async () => {
    const token = localStorage.getItem("token");
    try {
      const response = await axios.delete(
        `${process.env.API_ENDPOINT}recipe/${id}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        }
      );
      console.log(response);
      setSuccessMessage("Recipe deleted successfully");
      setErrorMessage(null);
    } catch (error) {
      console.log(error);
      setErrorMessage(error.response.data.error || "An error occurred");
      setSuccessMessage(null);
    }
  };

  return (
    <div className={classes.card}>
      <h2>Edit Recipe</h2>
      <form className={classes.form}>
        <div>
          <div className={classes.control}>
            <label htmlFor="name">Recipe Name:</label>
            <input
              type="text"
              id="name"
              value={name}
              onChange={(event) => setName(event.target.value)}
            />
          </div>
          <div className={classes.control}>
            <label htmlFor="image">Recipe Image:</label>
            <input
              type="file"
              id="image"
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
              value={category}
              onChange={(event) => setCategory(event.target.value)}
            >
              <option value="">Select Category</option>
              <option value="salads">Salads</option>
              <option value="soups">Soups</option>
              <option value="desserts">Desserts</option>
              <option value="main-dishes">Main Dishes</option>
              <option value="sides">Sides</option>
              <option value="breakfast">Breakfast</option>
              <option value="drinks">Drinks</option>
              <option value="snacks">Snacks</option>
            </select>
          </div>
          <div className={classes.control}>
            <label htmlFor="totalTime">Total Time:</label>
            <input
              type="number"
              id="totalTime"
              min="0"
              value={totalTime}
              onChange={(event) => setTotalTime(event.target.value)}
            />
          </div>
        </div>
        <div>
          <div className={classes.control}>
            <label htmlFor="ingredients">Ingredients:</label>
            <textarea
              id="ingredients"
              rows="5"
              placeholder="Insert (;) after each ingredients"
              value={ingredients}
              onChange={(event) => setIngredients(event.target.value)}
            ></textarea>
          </div>
          <div className={classes.control}>
            <label htmlFor="instructions">Instructions:</label>
            <textarea
              id="instructions"
              rows="10"
              value={instructions}
              onChange={(event) => setInstructions(event.target.value)}
            ></textarea>
          </div>
        </div>
        <div className={classes.buttons}>
          <div className={classes.actions_delete}>
            <button onClick={deleteRecipe}>Delete Recipe</button>
          </div>
          <div className={classes.actions}>
            <button onClick={submitHandler}>Update Recipe</button>
          </div>
        </div>
        {errorMessage && <div className={classes.error}>{errorMessage}</div>}
        {successMessage && (
          <div className={classes.success}>{successMessage}</div>
        )}
      </form>
    </div>
  );
};

export default EditRecipe;
