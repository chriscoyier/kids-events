import Head from "next/head";
import styles from "../styles/Home.module.scss";

export async function getServerSideProps() {
  const res = await fetch(
    `${process.env.SITE_URL}/.netlify/functions/portland`
  );
  const data = await res.json();

  return { props: { data } };
}

type kidEvent = {
  title: string;
  date: string;
  url: string;
  venue: string;
};

export default function Home(props: { data: kidEvent[] }) {
  return (
    <>
      <Head>
        <title>Kids Events</title>
        <meta name="description" content="Kids Events in Oregon" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link
          rel="icon"
          href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>👧</text></svg>"
        />
      </Head>

      <main className={styles.main}>
        <h1>Kids Events</h1>

        <div className={styles.grid}>
          {props.data.map((event) => {
            const date = new Date(event.date).toDateString();

            return (
              <div key={event.title} className={styles.card}>
                <dl className={styles.eventList}>
                  <dt>Title</dt>
                  <dd>
                    <h3>
                      <a href={event.url}>{event.title}</a>
                    </h3>
                  </dd>

                  <dt>Date</dt>
                  <dd>
                    <time>{date}</time>
                  </dd>

                  <dt>Venue</dt>
                  <dd>{event.venue}</dd>
                </dl>
              </div>
            );
          })}
        </div>
      </main>
    </>
  );
}
