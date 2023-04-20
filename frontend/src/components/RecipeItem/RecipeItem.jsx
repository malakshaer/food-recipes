import { useRouter } from "next/router";
import classes from "./RecipeItem.module.css";
import Image from "next/image";
import recipeImage from "../../../public/Spaghetti.jpg";
import LikeButton from "../LikeButton/LikeButton";
import SaveButton from "../SaveButton/SaveButton";
import { FaEdit } from "react-icons/fa";
import { useState } from "react";
import axios from "axios";

const RecipeItem = (props) => {
  const router = useRouter();
  const [saved, setSaved] = useState(false);

  const showRecipeDetails = () => {
    router.push(`/recipe-details/${props.id}`);
  };

  const handleEditRecipe = () => {
    router.push(`/edit-recipe/${props.id}`);
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

  const handleAction = async () => {
    const token = localStorage.getItem("token");
    try {
      const response = await axios({
        method: "post",
        url: `${process.env.API_ENDPOINT}${
          saved ? "unsave_recipe" : "save_recipe"
        }/${props.id}`,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      console.log(response);
      setSaved(!saved);
    } catch (error) {
      console.log(error);
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
        </div>
        <span>⏱{props.time} min</span>
        <div className={classes.user}>
          <Image
            src={`data:image/*;base64,${props.recipeAuthorImage}`}
            alt="author-profile-image"
            width={25}
            height={25}
          />
          <p>{props.authorName}</p>
        </div>
        <div className={classes.actions}>
          {props.showButton && (
            <div className={classes.buttonWrapper}>
              <button onClick={handleEditRecipe}>
                <FaEdit />
                Edit Recipe
              </button>
            </div>
          )}
          <div className={classes.buttonWrapper}>
            <SaveButton onClick={handleAction} saved={saved} />
          </div>
          <div className={classes.buttonWrapper}>
            <LikeButton />
            <span>{props.likes}</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RecipeItem;
