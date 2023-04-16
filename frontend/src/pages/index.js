import Head from "next/head";
import HomePage from "./HomePage/HomePage";

function Home() {
  return (
    <div>
      <Head>
        <title>Food Recipes</title>
        <meta name="description" content="A recipe sharing app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <HomePage />
      </main>
    </div>
  );
}

export default Home;
