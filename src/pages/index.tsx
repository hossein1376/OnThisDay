import Head from "next/head";
import { Inter } from "next/font/google";

import styles from "@/styles/Home.module.css";
import EventType from "./EventType";
import { useEffect, useState } from "react";
import {
  Checkbox,
  Col,
  Divider,
  InputNumber,
  Row,
  Skeleton,
  Slider,
} from "antd";
import type { CheckboxValueType } from "antd/es/checkbox/Group";

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

  const [selected, setSelected] = useState(true);
  const [events, setEvents] = useState(true);
  const [holidays, setHolidays] = useState(false);
  const [births, setBirths] = useState(false);
  const [deaths, setDeaths] = useState(false);

  const [yearRange, setYearRange] = useState<[number, number]>([0, 2023]);

  const date = new Date().toJSON().slice(0, 10).replace(/-/g, "/");
  useEffect(() => {
    const url = "https://docker.iran.liara.run/events";
    fetch(url)
      .then((res) => res.json())
      .then((data: Data) => {
        setData(data);
        setLoading(false);
      });
  }, []);

  const groupChangeHandler = (checkedValues: CheckboxValueType[]) => {
    setSelected(checkedValues.includes("Selected"));
    setEvents(checkedValues.includes("Events"));
    setHolidays(checkedValues.includes("Holidays"));
    setBirths(checkedValues.includes("Births"));
    setDeaths(checkedValues.includes("Deaths"));
  };

  const sliderChangeHandler = (values: [number, number]) => {
    setYearRange(values);
  };

  const startYearRangeChangeHandler = (value: number | null) => {
    if (value !== null) {
      setYearRange([value, yearRange[1]]);
    }
  };

  const endYearRangeChangeHandler = (value: number | null) => {
    if (value !== null) {
      setYearRange([yearRange[0], value]);
    }
  };

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
        {loading && (
          <div className={styles.boxBody}>
            <Skeleton paragraph={{ rows: 8 }} />
            <br />
            <br />
            <Skeleton paragraph={{ rows: 8 }} />
          </div>
        )}

        {!loading && (
          <>
            <div className={styles.head}>
              <h1 className={styles.heading}>On This Day:</h1>
              <p className={styles.subheading}>{date}</p>

              <Checkbox.Group
                defaultValue={["Selected", "Events"]}
                onChange={groupChangeHandler}
              >
                <Row>
                  <Col>
                    <Checkbox value="Selected" className={styles.checkboxLabel}>
                      Selected Events
                    </Checkbox>
                  </Col>
                  <Col>
                    <Checkbox value="Events" className={styles.checkboxLabel}>
                      Events
                    </Checkbox>
                  </Col>
                  <Col>
                    <Checkbox value="Holidays" className={styles.checkboxLabel}>
                      Holidays
                    </Checkbox>
                  </Col>
                  <Col>
                    <Checkbox value="Births" className={styles.checkboxLabel}>
                      Births
                    </Checkbox>
                  </Col>
                  <Col>
                    <Checkbox value="Deaths" className={styles.checkboxLabel}>
                      Deaths
                    </Checkbox>
                  </Col>
                </Row>
              </Checkbox.Group>

              <Divider />
              <p className={styles.yearRange}>Year Range:</p>

              <Row className={styles.sliderBox}>
                <InputNumber
                  min={0}
                  max={2023}
                  value={yearRange[0]}
                  onChange={(value) => startYearRangeChangeHandler(value)}
                />
                <Col style={{ width: "100%" }}>
                  <Slider
                    range={{ draggableTrack: true }}
                    min={0}
                    max={2023}
                    defaultValue={[0, 2023]}
                    onChange={sliderChangeHandler}
                  />
                </Col>
                <InputNumber
                  min={0}
                  max={2023}
                  value={yearRange[1]}
                  onChange={(value) => endYearRangeChangeHandler(value)}
                />
              </Row>
            </div>
            {selected && (
              <EventType
                title={"Selected Events"}
                data={data?.selected}
                yearRange={yearRange}
              ></EventType>
            )}

            {events && (
              <EventType
                title={"Events"}
                data={data?.events}
                yearRange={yearRange}
              ></EventType>
            )}

            {holidays && (
              <EventType
                title={"Holidays"}
                data={data?.holidays}
                yearRange={yearRange}
              ></EventType>
            )}

            {births && (
              <EventType
                title={"Births"}
                data={data?.births}
                yearRange={yearRange}
              ></EventType>
            )}

            {deaths && (
              <EventType
                title={"Deaths"}
                data={data?.deaths}
                yearRange={yearRange}
              ></EventType>
            )}
          </>
        )}
      </main>
    </>
  );
}
