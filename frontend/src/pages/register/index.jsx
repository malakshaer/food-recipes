import classes from "./Register.module.css";
import { useRef } from "react";
import Link from "next/link";

const Register = () => {
  const userNameInputRef = useRef();
  const emailInputRef = useRef();
  const passwordInputRef = useRef();
  const confirmPasswordInputRef = useRef();

  const handleSubmit = async (event) => {
    event.preventDefault();

    const userName = userNameInputRef.current.value;
    const email = emailInputRef.current.value;
    const password = passwordInputRef.current.value;
    const confirmPassword = confirmPasswordInputRef.current.value;

    const data = {
      userName,
      email,
      password,
      confirmPassword,
    };
    console.log(data);
  };

  return (
    <form className={classes.form} onSubmit={handleSubmit}>
      <h3>Register</h3>
      <div className={classes.input}>
        <input
          className={classes.form__input}
          type="text"
          id="userName"
          required
          placeholder="User Name"
          ref={userNameInputRef}
        />
      </div>
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
      <div className={classes.input}>
        <input
          className={classes.form__input}
          type="password"
          id="confirmPassword"
          required
          placeholder="Confirm Password"
          ref={confirmPasswordInputRef}
        />
      </div>
      <div className={classes.footer}>
        <button>Register</button>
        <div>
          <p>
            Already have an account <Link href="/login">Login</Link>
          </p>
        </div>
      </div>
    </form>
  );
};

export default Register;
