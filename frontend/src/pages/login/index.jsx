import classes from "./Login.module.css";
import { useRef } from "react";
import Link from "next/link";

const Login = () => {
  const emailInputRef = useRef();
  const passwordInputRef = useRef();

  const handleSubmit = async (event) => {
    event.preventDefault();

    const email = emailInputRef.current.value;
    const password = passwordInputRef.current.value;

    const data = {
      email,
      password,
    };
    console.log(data);
  };

  return (
    <form className={classes.form} onSubmit={handleSubmit}>
      <h3>Register</h3>

      <div className={classes.input}>
        <input
          type="email"
          id="email"
          required
          className={classes.form__input}
          placeholder="Email"
          ref={emailInputRef}
        />
      </div>
      <div className={classes.input}>
        <input
          className={classes.form__input}
          type="password"
          id="password"
          required
          placeholder="Password"
          ref={passwordInputRef}
        />
      </div>
      <div className={classes.footer}>
        <button>Login</button>
        <div>
          <p>
            Already have an account <Link href="/register">Register</Link>
          </p>
        </div>
      </div>
    </form>
  );
};

export default Login;
