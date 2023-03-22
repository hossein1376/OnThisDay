import styles from "@/styles/Home.module.css";

interface Props {
  title : string;
}

export default function EventType({title}: Props): JSX.Element {
  return (
    <div className={styles.box}>
      <h2 className={styles.boxHeader}>{title}</h2>
      <ul>
        <li>random stuff</li>
        <li>random stuff</li>
        <li>random stuff</li>
      </ul>
    </div>
  )
}