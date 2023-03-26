import styles from "@/styles/Home.module.css";
import { Image } from "antd";
import { useState } from "react";
import type { Events } from "./Types";

export interface Props {
  event: Events;
}

export default function EventDetails({ event }: Props): JSX.Element {
  const [details, setDetails] = useState(false);

  const date = new Date();

  return (
    <>
      <li className={styles.boxElement} key={event.year}>
        <div>
          {event.year !== 0 && <h3 className={styles.year}>{event.year}</h3>}
          {event.year !== 0 && (
            <h4 className={styles.yearsPast}>
              {date.getFullYear() - event.year} years ago
            </h4>
          )}

          <p className={styles.summery}>
            {event.text}&nbsp;
            <a
              href={event.pages[0].content_urls.desktop.page}
              target={"_blank"}
              className={styles.readMore}
            >
              Read more⧉
            </a>
          </p>

          <p
            className={styles.viewDetails}
            onClick={() => setDetails(!details)}
          >
            {!details ? "View Details ▼" : "Hide Details ▲"}
          </p>

          {details && (
            <p className={styles.details}>{event.pages[0].extract}</p>
          )}
        </div>
        {event.pages[0]?.thumbnail?.source &&
          event.pages[0].thumbnail.width && (
            <div className={styles.thumbnail}>
              <Image
                width={event.pages[0].thumbnail.width}
                src={event.pages[0].thumbnail.source}
                alt={event.pages[0].description}
              />
              <p className={styles.caption}>{event.pages[0].description}</p>
            </div>
          )}
      </li>
    </>
  );
}