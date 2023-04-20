import RecipeItem from "../RecipeItem/RecipeItem";
import classes from "./RecipeList.module.css";

const RecipeList = (props) => {
  const { recipes, savedRecipe } = props;

  if (!recipes || recipes.length === 0) {
    return (
      <div className={classes.no_recipes}>
        <div>No recipes yet.</div>
      </div>
    );
  }

  return (
    <ul className={classes.list}>
      {recipes.map((recipe) => (
        <RecipeItem
          key={recipe.id}
          id={recipe.id}
          name={recipe.name}
          ingredients={recipe.ingredients}
          instructions={recipe.instructions}
          image={recipe.recipeimage}
          time={recipe.totaltime}
          category={recipe.recipecategory}
          likes={recipe.likes}
          authorName={recipe.recipeauthorname}
          recipeAuthorImage={recipe.recipeauthorimage}
          showButton={props.showButton}
          savedRecipe={savedRecipe}
        />
      ))}
    </ul>
  );
};

export default RecipeList;
