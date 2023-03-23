import styles from "@/styles/Home.module.css";
import { Divider, Image, Space } from "antd";

interface Props {
  title: string;
  data: Events[] | undefined;
  yearRange: [number, number];
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

export default function EventType({
  title,
  data,
  yearRange,
}: Props): JSX.Element {
  const events = data?.map((event) => {
    if (event.year >= yearRange[0] && event.year <= yearRange[1]) {
      return (
        <>
          <li className={styles.boxElement} key={event.year}>
            <div>
              {event.year !== 0 && <h3>{event.year}</h3>}
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
    }
  });
  return (
    <Space direction="vertical" size="large" className={styles.box}>
      <h2 className={styles.boxHeader}>{title}</h2>
      <div className={styles.boxBody}>
        <ul>{events}</ul>
      </div>
    </Space>
  );
}
