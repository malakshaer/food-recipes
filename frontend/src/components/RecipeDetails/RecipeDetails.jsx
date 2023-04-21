import classes from "./RecipeDetails.module.css";
import Image from "next/image";
import { useRouter } from "next/router";
import recipeImage from "../../../public/Spaghetti.jpg";
import LikeButton from "../LikeButton/LikeButton";
import SaveButton from "../SaveButton/SaveButton";
import { useState, useEffect } from "react";
import axios from "axios";
import ShowProfile from "../ShowProfile/ShowProfile";

const RecipeDetails = (props) => {
  const router = useRouter();
  const [recipe, setRecipe] = useState({});
  const [saved, setSaved] = useState(false);
  const { id } = router.query;
  // const { id } = props;

  useEffect(() => {
    const fetchRecipeData = async () => {
      try {
        const token = localStorage.getItem("token");
        const authHeaders = {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        };
        const response = await axios.get(
          `${process.env.API_ENDPOINT}recipe/${id}`,
          {
            headers: authHeaders,
          }
        );
        const recipe = response.data.recipe;
        console.log(recipe);
        setRecipe(recipe);
      } catch (error) {
        console.log(error);
      }
    };
    fetchRecipeData();
  }, [props.id]);

  const handleSave = async () => {
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
  const createdDate = new Date(recipe.recipecreatedat);
  const formattedDate = createdDate.toLocaleString("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "numeric",
    minute: "numeric",
    // second: "numeric",
    // timeZoneName: "short",
  });

  return (
    <div className={classes.card}>
      <div className={classes.userInfo}>
        <Image
          src={`data:image/*;base64,${recipe.recipeauthorimage}`}
          alt="user-profile-image"
          width={50}
          height={50}
        />
        <div className={classes.userText}>
          <span>{recipe.recipeauthorname}</span>
          <d>{recipe.recipeauthorbio}</d>
        </div>
        <ShowProfile id={recipe.recipeauthorid} />
      </div>
      <div className={classes.container}>
        <div className={classes.left}>
          <div className={classes.header}>
            <h1>{recipe.name}</h1>
            <span>Category: {recipe.recipecategory}</span>
            <span>Time needed: 🕐{recipe.totaltime} min</span>
            <span>Date: {formattedDate}</span>
            <div className={classes.actions}>
              <SaveButton saved={props.saved} onClick={handleSave} />
              <LikeButton />
              <span>{recipe.likes}</span>
            </div>
          </div>
          <div className={classes.image}>
            <Image
              src={`data:image/*;base64,${recipe.recipeimage}`}
              alt="recipe-image"
              width={300}
              height={250}
            />
          </div>
        </div>
        <div className={classes.content}>
          {recipe && recipe.ingredients && (
            <div>
              <h3>Ingredients:</h3>
              <ul>
                {recipe.ingredients.map((ingredient, index) => (
                  <li key={index}>{ingredient.text}</li>
                ))}
              </ul>
            </div>
          )}
          <div>
            <h3>Instructions:</h3>
            <p>{recipe.instructions}</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export async function getServerSideProps(context) {
  const { contextId } = context.query;

  return {
    props: { contextId },
  };
}

export default RecipeDetails;
