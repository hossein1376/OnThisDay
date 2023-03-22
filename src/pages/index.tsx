import { useEffect } from "react";

import Head from "next/head";
import { Inter } from "next/font/google";

import styles from "@/styles/Home.module.css";
import EventType from "./EventType";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  useEffect(() => {
    // get data from API
    const url = "https://api.onthisday.com/events?date=2021-09-01";
    fetch(url)
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
      });
  }, []);

  return (
    <>
      <Head>
        <title>On This Day ...</title>
        <meta
          name="description"
          content="What happened on this day in history?"
        />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={styles.main}>
        <h1>On This Day:</h1>
        <EventType title={"this"}></EventType>
        <EventType title={"that"}></EventType>
      </main>
    </>
  );
}
