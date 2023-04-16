// our-domain.com/new-recipe

import CreateRecipe from "../components/CreateRecipe/CreateRecipe";

const NewRecipePage = () => {
  const addRecipeHandler = (recipeData) => {
    console.log(recipeData);
  };

  return <CreateRecipe onAddRecipe={addRecipeHandler} />;
};

export default NewRecipePage;
