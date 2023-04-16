import { useRouter } from "next/router";
import classes from "./RecipeItem.module.css";
import Image from "next/image";
import recipeImage from "../../../public/Spaghetti.jpg";

const RecipeItem = (props) => {
  const router = useRouter();

  const showRecipeDetails = () => {
    router.push("/" + props.id);
  };

  return (
    <div className={classes.card}>
      <div className={classes.image}>
        <Image src={recipeImage} alt="recipe-image" priority="high" />
      </div>
      <div className={classes.content}>
        <h1>{props.name}</h1>
        <p>({props.category})</p>
        <div className={classes.time}>
          <h3>⏱</h3>
          <p>{props.total_time} minutes</p>
        </div>
        <h3>Created At:</h3>
        <p>{props.created_at}</p>
      </div>
      <div>
        <button className={classes.actions} onClick={showRecipeDetails}>
          Show
        </button>
      </div>
    </div>
  );
};

export default RecipeItem;
