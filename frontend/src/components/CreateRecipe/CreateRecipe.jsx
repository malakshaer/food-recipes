import { useRef } from "react";
import classes from "./CreateRecipe.module.css";

const CreateRecipe = (props) => {
  const nameInputRef = useRef();
  const imageInputRef = useRef();
  const instructionsInputRef = useRef();
  const ingredientsInputRef = useRef();

  const submitHandler = (event) => {
    event.preventDefault();

    const enteredName = nameInputRef.current.value;
    const enteredImage = imageInputRef.current.value;
    const enteredInstructions = instructionsInputRef.current.value;
    const enteredIngredients = ingredientsInputRef.current.value;
    const ingredients = enteredIngredients.split(";");

    const recipeData = {
      name: enteredName,
      image: enteredImage,
      instructions: enteredInstructions,
      ingredients: ingredients,
    };

    props.onAddRecipe(recipeData);
  };

  return (
    <div className={classes.card}>
      <form className={classes.form} onSubmit={submitHandler}>
        <div className={classes.control}>
          <label htmlFor="title">Recipe Name:</label>
          <input type="text" required id="name" ref={nameInputRef} />
        </div>
        <div className={classes.control}>
          <label htmlFor="image">Recipe Image:</label>
          <input type="url" required id="image" ref={imageInputRef} />
        </div>
        <div className={classes.control}>
          <label htmlFor="ingredients">Ingredients</label>
          <textarea
            type="text"
            required
            id="ingredients"
            ref={ingredientsInputRef}
            placeholder="Insert (;) after each ingredients"
          />
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
      </form>
    </div>
  );
};

export default CreateRecipe;
