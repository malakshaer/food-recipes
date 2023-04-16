import classes from "./NavBar.module.css";
import Link from "next/link";

const NavBar = () => {
  return (
    <header className={classes.header}>
      <div className={classes.logo}>Food Recipes</div>
      <nav>
        <ul>
          <li>
            <Link href="/">All Recipes</Link>
          </li>
          <li>
            <Link href="/create-recipe">Add New Recipe</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default NavBar;
