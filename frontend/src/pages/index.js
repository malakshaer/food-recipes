import Head from "next/head";
import Register from "./register";

function Home() {
  return (
    <div>
      <Head>
        <title>Food Recipes</title>
        <meta name="description" content="A recipe sharing app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <Register />
      </main>
    </div>
  );
}

export default Home;
