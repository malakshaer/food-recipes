import { useRef } from "react";
import classes from "./EditRecipe.module.css";

const recipes = [
  {
    id: 1,
    name: "Spaghetti",
    ingredients: [
      { id: 1, text: "1/2 cup olive oil" },
      { id: 2, text: "1 onion, chopped" },
      { id: 3, text: "1 green bell pepper, chopped" },
      { id: 4, text: "2 cloves garlic, chopped" },
      { id: 5, text: "1 (28 ounce) can crushed tomatoes" },
      { id: 6, text: "1 (6 ounce) can tomato paste" },
      { id: 7, text: "1/2 cup water" },
    ],
    instructions:
      "Heat oil in a large pot over medium heat. Cook and stir onion, green bell pepper, and garlic in the hot oil until onion has softened and turned translucent, about 5 minutes. Stir crushed tomatoes, tomato paste, and water into the onion mixture; season with salt and pepper. Bring sauce to a boil, reduce heat to medium-low, and simmer until flavors have blended, about 30 minutes.",
    image: "",
    total_time: "10min",
    category: "meal",
    created_at: "1/1/2023",
  },
];
const EditRecipe = (props) => {
  const recipe = props.recipes
    ? props.recipes.find((recipe) => recipe.id === props.id)
    : null;

  const nameInputRef = useRef();
  const imageInputRef = useRef();
  const instructionsInputRef = useRef();
  const ingredientsInputRef = useRef();
  const categoryInputRef = useRef();
  const totalTimeInputRef = useRef();

  const nameDefaultValue = recipe ? recipe.name : "";
  const imageDefaultValue = recipe ? recipe.image : "";
  const categoryDefaultValue = recipe ? recipe.category : "";
  const totalTimeDefaultValue = recipe ? recipe.total_time : "";
  const ingredientsDefaultValue = recipe ? recipe.ingredients.join("; ") : "";

  const submitHandler = (event) => {
    event.preventDefault();

    const recipeData = {
      name: enteredName,
      image: enteredImage,
      instructions: enteredInstructions,
      ingredients: ingredients,
      category: enteredCategory,
      total_time: enteredTotalTime,
    };

    props.onEditRecipe(props.id, recipeData);
  };

  return (
    <div className={classes.card}>
      <h2>Edit Recipe</h2>
      <form className={classes.form} onSubmit={submitHandler}>
        <div>
          <div className={classes.control}>
            <label htmlFor="name">Recipe Name:</label>
            <input
              type="text"
              required
              id="name"
              defaultValue={nameDefaultValue}
              ref={nameInputRef}
            />
          </div>
          <div className={classes.control}>
            <label htmlFor="image">Recipe Image:</label>
            <input
              type="file"
              required
              id="image"
              ref={imageInputRef}
              defaultValue={imageDefaultValue}
            />
          </div>
          <div className={classes.control}>
            <label htmlFor="category">Category:</label>
            <select
              className={classes.category}
              id="category"
              required
              ref={categoryInputRef}
              defaultValue={categoryDefaultValue}
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
              required
              id="totalTime"
              defaultValue={totalTimeDefaultValue}
              ref={totalTimeInputRef}
              min="0"
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
              ref={ingredientsInputRef}
              defaultValue={ingredientsDefaultValue}
              placeholder="Insert (;) after each ingredients"
            ></textarea>
          </div>
          <div className={classes.control}>
            <label htmlFor="instructions">Instructions:</label>
            <textarea
              id="instructions"
              required
              rows="10"
              ref={instructionsInputRef}
              defaultValue={ingredientsDefaultValue}
            ></textarea>
          </div>
        </div>
        <div className={classes.buttons}>
          <div className={classes.actions_delete}>
            <button>Delete Recipe</button>
          </div>
          <div className={classes.actions}>
            <button>Update Recipe</button>
          </div>
        </div>
      </form>
    </div>
  );
};

export default EditRecipe;
