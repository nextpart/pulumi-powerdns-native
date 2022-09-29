import * as powerdns from "@pulumi/powerdns";

const random = new powerdns.Random("my-random", { length: 24 });

export const output = random.result;