import dayjs from "dayjs";

dayjs.locale("ja");

export const convertDate = (date: Date | string) => {
  return dayjs(date).format("YYYY-MM-DD HH:mm:ss");
};
