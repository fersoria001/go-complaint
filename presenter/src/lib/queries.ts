/* eslint-disable @typescript-eslint/no-explicit-any */

import {
  City,
  Complaint,
  ComplaintTypeList,
  Country,
  County,
  Employee,
  Enterprise,
  Industry,
  PhoneCode,
  Receiver,
  User,
  UserDescriptor,
  UsersForHiring,
} from "./types";

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export const CountriesQuery = (): string => `
    {
        Countries {
        id
        name
        }
    }
`;
export const CountryListType = (data: any): Country[] => {
  return data.Countries as Country[];
};
export const CountiesQuery = (id: number): string => `
  {
    Counties(ID: ${id}) {
      id
      name
    }
  }    
`;
export const CountyListType = (data: any): County[] => {
  return data.Counties as County[];
};
export const CitiesQuery = (id: number): string => `
  {
    Cities(ID: ${id}) {
      id
      name
    }
  }
`;
export const CityListType = (data: any): City[] => {
  return data.Cities as City[];
};
export const PhoneCodeQuery = (id: number): string => `
  {
    PhoneCode(ID: ${id}) {
      id
      code
    }
  }
`;
export const PhoneCodeType = (data: any): PhoneCode => {
  return data.PhoneCode as PhoneCode;
};
export const SignInQuery = (
  email: string,
  password: string,
  rememberMe: boolean
): string => `
   {
    Login(email: "${email}", password: "${password}", rememberMe: ${rememberMe}) {
      token
    }
  }
`;
export const SignIn = (data: any): string => {
  console.log(data);
  return data.Login.token;
};
export const UserDescriptorQuery = (): string => `
  {
    UserDescriptor {
      email
      fullName
      profileIMG
      ip
    }
  }
`;
export const UserDescriptorType = (data: any): UserDescriptor => {
  return data.UserDescriptor as UserDescriptor;
};
export const IndustriesQuery = (): string => `
  {
    Industries {
      id
      name
    }
  }
`;
export const IndustryListType = (data: any): Industry[] => {
  return data.Industries as Industry[];
};
export const IsEnterpriseNameAvailableQuery = (name: string): string => `
  {
    IsEnterpriseNameAvailable(ID: "${name}")
  }
`;
export const IsEnterpriseNameAvailable = (data: any): boolean => {
  return data.IsEnterpriseNameAvailable;
};

export const OwnerEnterprisesQuery = (): string => `
  {
    OwnerEnterprises {
      name
      bannerIMG
      logoIMG
      email
      website
      phone
      industry
      address { country county city }
      foundationDate
    }
  }
`;
export const OwnerEnterprisesTypeList = (data: any): Enterprise[] => {
  return data.OwnerEnterprises;
};
export const FindReceiverQuery = (name: string): string => `
  {
    FindReceiver(term: "${name}") {
      ID
      fullName
      IMG
    }
  }
`;

export const ReceiverTypeList = (data: any): Receiver[] => {
  return data.FindReceiver;
};

export const IsValidReceiverQuery = (name: string): string => `
  {
    IsValidReceiver(ID: "${name}")
  }
`;
export const IsValidReceiver = (data: any): boolean => {
  return data.IsValidReceiver;
};

export const DraftQuery = (
  id: string,
  limit: number,
  offset: number
): string => `
  {
    Draft(ID: "${id}", limit:${limit}, offset:${offset}) {
      complaints {
      id
      authorFullName
      authorProfileIMG
      receiverFullName
      receiverProfileIMG
      status
      message { title description body }
      createdAt
      updatedAt
      }
      count
      currentLimit
      currentOffset
    }
  }
`;

export const DraftTypeList = (data: any): ComplaintTypeList => {
  return data.Draft as ComplaintTypeList;
};

export const SentQuery = (
  id: string,
  limit: number,
  offset: number
): string => `
  {
    Sent(ID: "${id}", limit:${limit}, offset:${offset}) {
      complaints {
      id
      authorFullName
      authorProfileIMG
      receiverFullName
      receiverProfileIMG
      status
      message { title description body }
      createdAt
      updatedAt
      }
      count
      currentLimit
      currentOffset
    }
  }
`;
export const SentTypeList = (data: any): ComplaintTypeList => {
  return data.Sent as ComplaintTypeList;
};
export const ComplaintQuery = (id: string): string => `
  {
    Complaint(ID: "${id}") {
      authorID
      authorFullName
      authorProfileIMG
      receiverID
      receiverFullName
      receiverProfileIMG
      status
      message { title description body }
      createdAt
      updatedAt
      rating { rate comment }
    }
  }
`;

export const ComplaintType = (data: any): Complaint[] => {
  console.log("json data", data);
  return data.Complaint as Complaint[];
};

export const EnterpriseQuery = (id: string): string => `
  {
    Enterprise(ID: "${id}") {
      name
      bannerIMG
      logoIMG
      email
      website
      phone
      industry
      address { country county city }
      foundationDate
    }
  }
`;

export const EnterpriseType = (data: any): Enterprise => {
  return data.Enterprise as Enterprise;
};

export const UsersForHiringQuery = (
  id: string,
  limit: number,
  offset: number,
  query: string
): string => `
  {
    UsersForHiring(ID: "${id}", limit:${limit}, offset:${offset}, query: "${query}") {
      users{
      profileIMG
      email
      firstName
      lastName
      age
      phone
      address { country county city } 
      }
      count
      currentLimit
      currentOffset
    }
  }
`;

export const UsersForHiringType = (data: any): UsersForHiring => {
  return data.UsersForHiring as UsersForHiring;
};

export const UserQuery = (id: string): string => `
  {
    User(ID: "${id}") {
      profileIMG
      email
      firstName
      lastName
      age
      phone
      address { country county city }
    }
  }
`;

export const UserType = (data: any): User => {
  return data.User as User;
};

export const EmployeeQuery = (id: string): string => `
  {
    Employee(ID: "${id}") {
      ID
      profileIMG
      firstName
      lastName
      age
      email
      phone
      hiringDate
      approvedHiring
      approvedHiringAt
      position
    }
  }
`;

export const EmployeeType = (data: any): Employee => {
  return data.Employee as Employee;
};

export const EmployeesQuery = (id: string): string => `
  {
    Employees(ID: "${id}") {
      ID
      profileIMG
      firstName
      lastName
      age
      email
      phone
      hiringDate
      approvedHiring
      approvedHiringAt
      position
    }
  }
`;

export const EmployeesTypeList = (data: any): Employee[] => {
  console.log(data);
  return data.Employees as Employee[];
};

export async function Query<T>(
  queryFn: (...args: any[]) => string,
  castToFn: (data: any) => T,
  args: any[] = []
): Promise<T> {
  console.log(queryFn(...args));
  return fetch(import.meta.env.VITE_GRAPHQL_ENDPOINT, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify({ query: deleteLinebreaks(queryFn(...args)) }),
  })
    .then((res) => res.json())
    .then((data) => castToFn(data.data));
}

const deleteLinebreaks = (str: string): string => {
  return str.replace(/\n/g, "");
};
