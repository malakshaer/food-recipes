import classes from "./RecipeItem.module.css";

const RecipeItem = (props) => {
  return (
    <li className={classes.item}>
      <div className={classes.card}>
        <div className={classes.image}>
          <img src={props.image} alt={props.title} />
        </div>
        <div className={classes.content}>
          <h3>{props.name}</h3>
          <ul>
            {props.ingredients.map((ingredient, index) => (
              <li key={index}>{ingredient.text}</li>
            ))}
          </ul>
          <p>{props.instructions}</p>
        </div>
        <div>
          <button className={classes.actions}>Show</button>
        </div>
      </div>
    </li>
  );
};

export default RecipeItem;
