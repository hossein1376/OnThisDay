import styles from "@/styles/Home.module.css";
import { Divider, Image } from "antd";

interface Props {
  title: string;
  data: Events[] | undefined;
}

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

export default function EventType({ title, data }: Props): JSX.Element {
  const events = data?.map((event) => {
    return (
      <>
        <li className={styles.boxElement} key={event.year}>
          <div>
            <h3>{event.year}</h3>
            <p>{event.text}</p>
          </div>
          {event.pages[0]?.thumbnail?.source &&
            event.pages[0].thumbnail.width && (
              <Image
                width={event.pages[0].thumbnail.width}
                src={event.pages[0].thumbnail.source}
                alt={event.pages[0].title}
              />
            )}
        </li>
        <Divider />
      </>
    );
  });
  return (
    <div className={styles.box}>
      <h2 className={styles.boxHeader}>{title}</h2>
      <div className={styles.boxBody}>
        <ul>{events}</ul>
      </div>
    </div>
  );
}
