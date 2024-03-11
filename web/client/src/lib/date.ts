import dayjs, { locale } from "dayjs";

locale("ja");

export const convertDate = (date: Date | string) => {
  return dayjs(date).format("YYYY-MM-DD HH:mm:ss");
};
