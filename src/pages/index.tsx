import Head from "next/head";
import { Inter } from "next/font/google";

import styles from "@/styles/Home.module.css";
import EventType from "./EventType";
import { useEffect, useState } from "react";
import { Skeleton } from "antd";

const inter = Inter({ subsets: ["latin"] });

type Data = {
  selected: Events[];
  events: Events[];
  holidays: Events[];
  births: Events[];
  deaths: Events[];
};

type Events = {
  text: string;
  pages: Page[];
  year: number;
};

type Image = {
  source: string;
  width: number;
  height: number;
};

type Page = {
  title: string;
  pageid: number;
  thumbnail: Image;
  originalimage: Image;
  timestamp: string;
  description: string;
  content_urls: {
    desktop: {
      page: string;
    };
    mobile: {
      page: string;
    };
  };
  extract: string;
};

export default function Home() {
  const [data, setData] = useState<Data>();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const url = "https://docker.iran.liara.run/events";
    fetch(url)
      .then((res) => res.json())
      .then((data: Data) => {
        setData(data);
        setLoading(false);
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
        {loading && (
          <div className={styles.boxBody}>
            <Skeleton paragraph={{ rows: 7 }} />
            <br />
            <br />
            <Skeleton paragraph={{ rows: 7 }} />
          </div>
        )}

        {!loading && (
          <>
            <EventType
              title={"Selected Events"}
              data={data?.selected}
            ></EventType>

            <EventType title={"Events"} data={data?.events}></EventType>

            <EventType title={"Holidays"} data={data?.holidays}></EventType>

            <EventType title={"Births"} data={data?.births}></EventType>

            <EventType title={"Deaths"} data={data?.deaths}></EventType>
          </>
        )}
      </main>
    </>
  );
}
