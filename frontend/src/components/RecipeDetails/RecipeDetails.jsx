import classes from "./RecipeDetails.module.css";
import Image from "next/image";
import { useRouter } from "next/router";
import recipeImage from "../../../public/Spaghetti.jpg";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSave } from "@fortawesome/free-solid-svg-icons";
import LikeButton from "../LikeButton/LikeButton";
import SaveButton from "../SaveButton/SaveButton";

const RecipeDetails = (props) => {
  const router = useRouter();

  const showUserProfile = () => {
    router.push("/user-profile");
  };

  return (
    <div className={classes.card}>
      <div className={classes.userInfo}>
        <Image src={recipeImage} alt="user-profile-image" />
        <div className={classes.userText}>
          <span>Malak Shaer</span>
          <d>Master Chef</d>
        </div>
        <div className={classes.userInfoButton}>
          <button onClick={showUserProfile}>Show Profile</button>
        </div>
      </div>
      <div className={classes.container}>
        <div className={classes.left}>
          <div className={classes.header}>
            <h1>{props.name}</h1>
            <span>Category: {props.category}</span>
            <span>Time needed: 🕐{props.total_time} min</span>
            <span>Date: {props.created_at}</span>
            <div className={classes.actions}>
              <SaveButton />
              <LikeButton />
              <span>{props.likes}1</span>
            </div>
          </div>
          <div className={classes.image}>
            <Image src={recipeImage} alt="recipe-image" />
          </div>
        </div>
        <div className={classes.content}>
          <div>
            <h3>Ingredients:</h3>
            <ul>
              {props.ingredients.map((ingredient, index) => (
                <li key={index}>{ingredient.text}</li>
              ))}
            </ul>
          </div>
          <div>
            <h3>Instructions:</h3>
            <p>{props.instructions}</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RecipeDetails;
