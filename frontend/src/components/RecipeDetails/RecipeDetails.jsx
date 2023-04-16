import classes from "./RecipeDetails.module.css";
import Image from "next/image";
import recipeImage from "../../../public/Spaghetti.jpg";

const RecipeItem = (props) => {
  return (
    <div className={classes.card}>
      <div className={classes.header}>
        <div>
          <h1>{props.name}</h1>
          <p>({props.category})</p>
        </div>
        <div>
          <h3>Total Time:</h3>
          <p>{props.total_time} minutes</p>
          <h3>Created At:</h3>
          <p>{props.created_at}</p>
        </div>
      </div>
      <div className={classes.image}>
        <Image src={recipeImage} alt="recipe-image" priority="high" />
      </div>
      <div className={classes.content}>
        <h3>Ingredients:</h3>
        <ul>
          {props.ingredients.map((ingredient, index) => (
            <li key={index}>{ingredient.text}</li>
          ))}
        </ul>
        <h3>Instructions:</h3>
        <p>{props.instructions}</p>
      </div>
    </div>
  );
};

export default RecipeItem;
