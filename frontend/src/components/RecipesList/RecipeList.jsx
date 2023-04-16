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
          image={recipe.image}
          time={recipe.total_time}
          category={recipe.category}
          date={recipe.created_at}
        />
      ))}
    </ul>
  );
};

export default RecipeList;
