import { useRouter } from "next/router";
import classes from "./RecipeItem.module.css";
import Image from "next/image";
import recipeImage from "../../../public/Spaghetti.jpg";
import LikeButton from "../LikeButton/LikeButton";
import SaveButton from "../SaveButton/SaveButton";

const RecipeItem = (props) => {
  const router = useRouter();

  const showRecipeDetails = () => {
    router.push("/" + props.id);
  };

  const handleCardClick = (event) => {
    if (
      event.target.tagName !== "BUTTON" &&
      event.target.tagName !== "path" &&
      event.target.tagName !== "svg"
    ) {
      showRecipeDetails();
    }
  };

  return (
    <div className={classes.card} onClick={handleCardClick}>
      <div className={classes.image}>
        <Image src={recipeImage} alt="recipe-image" />
      </div>
      <div className={classes.content}>
        <div className={classes.time}>
          <h1>{props.name}</h1>
          <p>⏱{props.total_time} minutes</p>
        </div>
        <div className={classes.user}>
          <Image src={recipeImage} alt="author-profile-image" />
          <p>Malak Shaer</p>
        </div>
        <div className={classes.actions}>
          <div className={classes.buttonWrapper}>
            <SaveButton />
          </div>
          <div className={classes.buttonWrapper}>
            <LikeButton />
            <span>{props.likes}1</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RecipeItem;
