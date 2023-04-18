import RecipeItem from "../RecipeItem/RecipeItem";
import classes from "./RecipeList.module.css";

const RecipeList = (props) => {
  return (
    <ul className={classes.list}>
      {props.recipes.map((recipe) => (
        <RecipeItem
          key={recipe.id}
          id={recipe.id}
          name={recipe.name}
          ingredients={recipe.ingredients}
          instructions={recipe.instructions}
          image={recipe.recipeimage}
          time={recipe.totaltime}
          category={recipe.recipecategory}
          date={recipe.recipecreatedat}
          likes={recipe.likes}
          authorName={recipe.recipeauthorname}
          showButton={props.showButton}
        />
      ))}
    </ul>
  );
};

export default RecipeList;
