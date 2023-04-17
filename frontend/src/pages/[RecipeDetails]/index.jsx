import React from "react";
import RecipeDetails from "../../components/RecipeDetails/RecipeDetails";

const recipe = {
  id: 1,
  name: "Spaghetti",
  ingredients: [
    { id: 1, text: "1/2 cup olive oil" },
    { id: 2, text: "1 onion, chopped" },
    { id: 3, text: "1 green bell pepper, chopped" },
    { id: 4, text: "2 cloves garlic, chopped" },
    { id: 5, text: "1 (28 ounce) can crushed tomatoes" },
    { id: 6, text: "1 (6 ounce) can tomato paste" },
    { id: 7, text: "1/2 cup water" },
  ],
  instructions:
    "Heat oil in a large pot over medium heat. Cook and stir onion, green bell pepper, and garlic in the hot oil until onion has softened and turned translucent, about 5 minutes. Stir crushed tomatoes, tomato paste, and water into the onion mixture; season with salt and pepper. Bring sauce to a boil, reduce heat to medium-low, and simmer until flavors have blended, about 30 minutes.",
  image: "",
};

const index = () => {
  return (
    <RecipeDetails
      key={recipe.id}
      id={recipe.id}
      name={recipe.name}
      ingredients={recipe.ingredients}
      instructions={recipe.instructions}
      image={recipe.image}
    />
  );
};

export default index;
