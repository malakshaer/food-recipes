import { Fragment } from "react";
import RecipeList from "../../components/RecipesList/RecipeList";
import SearchBar from "../../components/SearchBar/SearchBar";

const FAKE_RECIPES = [
  {
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
    total_time: "10min",
    category: "meal",
    created_at: "1/1/2023",
  },
  {
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
    total_time: "10min",
    category: "meal",
    created_at: "1/1/2023",
  },
  {
    id: 2,
    name: "Onion Pie",
    ingredients: [
      { id: 1, text: "1/2 cup butter" },
      { id: 2, text: "1 onion, sliced" },
      { id: 3, text: "1 (10.75 ounce) can condensed cream of mushroom soup" },
      { id: 4, text: "1 (8 ounce) package cream cheese, softened" },
      { id: 5, text: "1 (8 ounce) container sour cream" },
      {
        id: 6,
        text: "1 (16 ounce) package frozen chopped spinach, thawed and squeezed dry",
      },
      { id: 7, text: "1 (9 inch) unbaked pie crust" },
    ],
    instructions:
      "Preheat oven to 350 degrees F (175 degrees C). Melt butter in a skillet over medium heat; cook and stir onion until tender and translucent, about 5 minutes. Mix onion, soup, cream cheese, sour cream, and spinach together in a bowl. Pour mixture into the pie crust. Bake in the preheated oven until set, 45 to 50 minutes.",
    image: "",
    total_time: "10min",
    category: "meal",
    created_at: "1/1/2023",
  },
];

const HomePage = () => {
  return (
    <Fragment>
      <SearchBar />
      <RecipeList recipes={FAKE_RECIPES} showButton={false} />
    </Fragment>
  );
};

export default HomePage;
