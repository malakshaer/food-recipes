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

  const handleEditRecipe = (id) => {
    router.push(`/edit-recipe/${id}`);
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

  const handleSave = async (id) => {
    const token = localStorage.getItem("token");
    try {
      const response = await axios({
        method: "post",
        url: `${process.env.API_ENDPOINT}save_recipe/${id}`,
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

  const handleLike = async (id) => {
    const token = localStorage.getItem("token");
    try {
      const response = await axios({
        method: "post",
        url: `${process.env.API_ENDPOINT}like_recipe/${id}`,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      console.log(response);
    } catch (error) {
      console.log(error);
    }
  };

  const handleUnSave = async (id) => {
    const token = localStorage.getItem("token");
    try {
      const response = await axios({
        method: "post",
        url: `${process.env.API_ENDPOINT}unsave_recipe/${id}`,
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
              <button onClick={() => handleEditRecipe(props.id)}>
                <FaEdit />
                Edit Recipe
              </button>
            </div>
          )}
          {props.showActionButton && (
            <div className={classes.like_save}>
              <div className={classes.buttonWrapper}>
                <SaveButton onClick={() => handleSave(props.id)} />
              </div>
              <div className={classes.buttonWrapper}>
                <LikeButton onClick={() => handleLike(props.id)} />
                <span>{props.likes}</span>
              </div>
            </div>
          )}
          {props.showUnActionButton && (
            <div className={classes.like_save}>
              <div className={classes.buttonWrapper}>
                <SaveButton onClick={() => handleUnSave(props.id)} />
              </div>
              <div className={classes.buttonWrapper}>
                <LikeButton onClick={() => handleLike(props.id)} />
                <span>{props.likes}</span>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default RecipeItem;
