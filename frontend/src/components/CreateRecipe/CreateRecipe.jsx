import { useRef } from "react";
import classes from "./CreateRecipe.module.css";

const CreateRecipe = (props) => {
  const nameInputRef = useRef();
  const imageInputRef = useRef();
  const instructionsInputRef = useRef();
  const ingredientsInputRef = useRef();
  const categoryInputRef = useRef();
  const totalTimeInputRef = useRef();

  const submitHandler = (event) => {
    event.preventDefault();

    const enteredName = nameInputRef.current.value;
    const enteredImage = imageInputRef.current.value;
    const enteredInstructions = instructionsInputRef.current.value;
    const enteredIngredients = ingredientsInputRef.current.value;
    const enteredCategory = categoryInputRef.current.value;
    const enteredTotalTime = totalTimeInputRef.current.value;
    const ingredients = enteredIngredients.split(";");

    const recipeData = {
      name: enteredName,
      image: enteredImage,
      instructions: enteredInstructions,
      ingredients: ingredients,
      category: enteredCategory,
      total_time: enteredTotalTime,
    };

    props.onAddRecipe(recipeData);
  };

  return (
    <div className={classes.card}>
      <form className={classes.form} onSubmit={submitHandler}>
        <h2>Unleash your creativity and let your taste buds dance with joy</h2>
        <div>
          <div className={classes.control}>
            <label htmlFor="name">Recipe Name:</label>
            <input type="text" required id="name" ref={nameInputRef} />
          </div>
          <div className={classes.control}>
            <label htmlFor="image">Recipe Image:</label>
            <input type="file" required id="image" ref={imageInputRef} />
          </div>
          <div className={classes.control}>
            <label htmlFor="category">Category:</label>
            <select
              className={classes.category}
              id="category"
              required
              ref={categoryInputRef}
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
              placeholder="Insert (;) after each ingredients"
            ></textarea>
          </div>
          <div className={classes.control}>
            <label htmlFor="instructions">Instructions:</label>
            <textarea
              id="instructions"
              required
              rows="5"
              ref={instructionsInputRef}
            ></textarea>
          </div>
          <div className={classes.actions}>
            <button>Add Recipe</button>
          </div>
        </div>
      </form>
    </div>
  );
};

export default CreateRecipe;
