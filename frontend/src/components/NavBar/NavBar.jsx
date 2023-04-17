import classes from "./NavBar.module.css";
import Link from "next/link";

const NavBar = () => {
  return (
    <header className={classes.header}>
      <div className={classes.logo}>FoodRecipes</div>
      <nav>
        <ul>
          <li>
            <Link href="/home">Home</Link>
          </li>
          <li>
            <Link href="/profile">Profile</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default NavBar;
