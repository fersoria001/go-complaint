import { dateFromMsString } from "@/lib/dateFromMsString";

export type TimeData = {
  occurredOn: string;
  qtty?: number;
};

export enum TimeDataBias {
  Day = "Day",
  Month = "Month",
}

export type ChartBuilderOptions = {
  day?: number;
  month?: number;
};

export type SelectorsData = {
  selectorLabel: string;
  selectors: Array<number>;
  defaultValue: number;
  xLabel: string;
};

export class ChartBuilder {
  private bias: TimeDataBias;
  private today: Date;
  private day: number;
  private month: number;
  private _selectorsData: SelectorsData;
  private ONE_DAY = 24 * 60 * 60 * 1000;

  constructor(bias: TimeDataBias, options?: ChartBuilderOptions) {
    this.today = new Date();
    this.bias = bias;
    this.day = options?.day
      ? options.day
      : parseInt(this.today.toUTCString().slice(5, 7));
    this.month = options?.month ? options.month : this.today.getMonth() + 1;
    switch (bias) {
      case TimeDataBias.Day:
        this._selectorsData = {
          selectorLabel: "day",
          selectors: this.daysOfThisMonth(),
          defaultValue: parseInt(this.today.toUTCString().slice(5, 7)),
          xLabel: "hours",
        };
        break;
      case TimeDataBias.Month:
        this._selectorsData = {
          selectorLabel: "month",
          selectors: this.monthsOfTheYear(),
          defaultValue: this.today.getMonth() + 1,
          xLabel: "days",
        };
        break;
    }
  }

  get selectorsData(): SelectorsData {
    return { ...this._selectorsData };
  }

  private monthsOfTheYear(): number[] {
    return [...Array(12)].map((_, i) => i + 1);
  }

  private daysOfThisMonth(): number[] {
    const monthFirst = new Date(this.today.getFullYear(), this.month - 1, 1);
    const nextMonthFirst = new Date(
      this.today.getFullYear(),
      this.month - 1 + 1,
      1
    );
    const qtty =
      (nextMonthFirst.getTime() - monthFirst.getTime()) / this.ONE_DAY;
    return [...Array(qtty)].map((_, i) => i + 1);
  }

  private monthly(data: TimeData[]): TimeData[] {
    const groups = Object.groupBy(data, ({ occurredOn }) =>
      dateFromMsString(occurredOn).toUTCString().slice(0, 16)
    );
    const keys = Object.keys(groups);
    const fillDates = this.daysOfThisMonth().map((v, i) => {
      const date = new Date(this.today.getFullYear(), this.month - 1, v);
      const stringDate = date.toUTCString().slice(0, 16);
      return {
        occurredOn: `${v}`,
        qtty: keys.includes(stringDate)
          ? groups[stringDate as keyof typeof groups]!.length
          : 0,
      };
    });
    return fillDates;
  }
  private daily(data: TimeData[]): TimeData[] {
    const todayDayNumber = new Date(
      this.today.getFullYear(),
      this.today.getMonth(),
      this.day
    )
      .toUTCString()
      .slice(0, 16);
    const data1 = data.filter(
      (v) =>
        dateFromMsString(v.occurredOn).toUTCString().slice(0, 16) ===
        todayDayNumber
    );
    const groups = Object.groupBy(data1, ({ occurredOn }) =>
      dateFromMsString(occurredOn).toUTCString().slice(17, 19)
    );
    const keys = Object.keys(groups);
    let fillDates = [...Array(24)].map((_, i) => {
      return {
        occurredOn: i.toString(),
        qtty: keys.includes(i.toString())
          ? groups[i.toString() as keyof typeof groups]!.length
          : 0,
      };
    });
    return fillDates;
  }

  filter(data: TimeData[]): TimeData[] {
    switch (this.bias) {
      case TimeDataBias.Day:
        return this.daily(data);
      case TimeDataBias.Month:
        return this.monthly(data);
      default:
        throw new Error("unexpected bias in ChartBuilder.filter");
    }
  }
}
