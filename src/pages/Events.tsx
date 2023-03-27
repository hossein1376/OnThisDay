import styles from "@/styles/Home.module.css";
import {Space} from "antd";
import type {Events} from "@/Types";
import EventDetails from "./EventDetails";
import useWindowDimensions from "@/hooks/getWindowDimensions";

interface Props {
  title: string;
  data: Events[] | undefined;
  yearRange: [number, number];
}

export default function Events({title, data, yearRange}: Props): JSX.Element {
  const {height, width} = useWindowDimensions();
  const size = width > 720 ? "large" : "small";

  const events = data?.map((event) => {
    if (event.year >= yearRange[0] && event.year <= yearRange[1]) {
      return (
        <ul key={event.pages[0].pageid}>
          <EventDetails event={event}/>
        </ul>
      );
    }
  });
  return (
    <>
      <Space direction="vertical" size={size} className={styles.box}>
        <h2 className={styles.boxHeader}>{title}</h2>
        <div className={styles.boxBody}>
          <ul>{events}</ul>
        </div>
      </Space>
    </>
  );
}
