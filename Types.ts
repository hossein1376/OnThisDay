export type Data = {
  selected: Events[];
  events: Events[];
  holidays: Events[];
  births: Events[];
  deaths: Events[];
};

export type Events = {
  text: string;
  pages: Page[];
  year: number;
};

export type Images = {
  source: string;
  width: number;
  height: number;
};

export type Page = {
  title: string;
  pageid: number;
  thumbnail: Images;
  originalimage: Images;
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