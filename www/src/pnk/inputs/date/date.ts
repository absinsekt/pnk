export const WeekdaysShort = ['пн', 'вт', 'ср', 'чт', 'пт', 'сб', 'вс'];
export const WeekdaysFull = ['понедельник', 'вторник', 'среда', 'четверг', 'пятница', 'суббота', 'воскресенье'];
export const MonthsFull = ['январь', 'февраль', 'март', 'апрель', 'май', 'июнь', 'июль', 'август', 'сентябрь', 'октябрь', 'ноябрь', 'декабрь'];

export type CalendarDay = {
  date: number;
  month: string;
  rawDate: Date;
  isToday: boolean;
  isSelected: boolean;
  isActiveMonth: boolean;
  isDisabled: boolean;
  isFirst: boolean;
  isLast: boolean;
}

export class CalendarConfig {
  constructor(
    public today = new Date(),
    public selectedDate: Date = null,
    public offsetDate: Date = null,
    public minDate: Date = null,
    public maxDate: Date = null,
    public isSundayFirst = false,
    public isWeekendDisabled = false,
  ) {}
}

const enum DateEquality {
  Day = 0,
  Month = 1,
  Year = 2,
}

function getStartOfMonth(now: Date): Date {
  const result = new Date(now);

  result.setDate(1);

  return result;
}

function getDayOffset(date: Date, isSundayFirst: boolean = false): number {
  let result = isSundayFirst
    ? date.getDay()
    : date.getDay() - 1;

  if (result === -1) {
    result = 6
  }

  return result;
}

function isEqualDate(date1: Date, date2: Date, eqType:DateEquality = DateEquality.Day): boolean {
  if (date1.getFullYear() !== date2.getFullYear()) return false;
  if (eqType === DateEquality.Year) return true;

  if (date1.getMonth() !== date2.getMonth()) return false;
  if (eqType === DateEquality.Month) return true;

  return date1.getDate() === date2.getDate();
}

export function getOffsetDate(date: Date, days: number, months: number, years: number) {
  const result = new Date(date);

  result.setDate(result.getDate() + days);
  result.setMonth(result.getMonth() + months);
  result.setFullYear(result.getFullYear() + years);

  return result;
}

export function getMonthCalendar(cfg = new CalendarConfig()): CalendarDay[] {
  const min = cfg.minDate.getTime();
  const max = cfg.maxDate.getTime();

  const result: CalendarDay[] = [];
  const position = getStartOfMonth(cfg.offsetDate);
  const offset = getDayOffset(position, cfg.isSundayFirst);

  position.setDate(-offset + 1);

  for (let i = 0; i < 42; i++) {
    const dt = new Date(position)
    const dtMills = dt.getTime();
    const dtDay = dt.getDay();
    const dtDate = dt.getDate();
    const dtMonth = dt.getMonth();

    result.push({
      date: dtDate,
      month: dtMonth.toString(), //TODO
      rawDate: dt,
      isToday: isEqualDate(dt, cfg.today),
      isSelected: cfg.selectedDate && isEqualDate(dt, cfg.selectedDate),
      isActiveMonth: isEqualDate(dt, cfg.offsetDate, DateEquality.Month),
      isDisabled: dtMills < min
        || dtMills > max
        || (cfg.isWeekendDisabled && (dtDay === 0 || dtDay === 6)),
      isFirst: isEqualDate(dt, cfg.minDate),
      isLast: isEqualDate(dt, cfg.maxDate),
    });

    position.setDate(position.getDate() + 1);
  }

  return result;
}
