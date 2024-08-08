export default function Home() {
  return (
      <article className="pt-8 lg:pt-16 mb-4 2xl:mb-8 lg:ps-5 flex flex-col">
        <p className="text-gray-700 text-sm md:text-md xl:text-xl 2xl:text-2xl 3xl:text-3xl mb-4 px-2">
          We have experienced issues related to our uptime, the explanation for this is that
          our cloud services were attacked using bots.
        </p>
        <p className="text-gray-700 text-sm md:text-md xl:text-xl 2xl:text-2xl 3xl:text-3xl mb-4 px-2">
          Bots are little programs that can be used for many purposes
          including process automation or interaction with users or other systems.
          When used to attack other systems, the attacker will start many of these little programs to
          constantly interact with the webpage in pursuit of the objective of taking it down.
        </p>
        <p className="text-gray-700 text-sm md:text-md xl:text-xl 2xl:text-2xl 3xl:text-3xl mb-4 px-2">
          Lucky for us, they didn&apos;t manage to take down our site but the constant requests
          that involved database operations incurred in high and unexpected maintenance costs.
          Go Complaint had to be temporarily shut down to avoid incurring in more unexpected costs until we manage to solve
          the situation.
        </p>
        <p className="text-gray-700 text-sm md:text-md xl:text-xl 2xl:text-2xl 3xl:text-3xl mb-4 px-2">
          In this downtime we are also improving our services and graphic user interface to provide better experience as
          soon as we can.
        </p>
        <p className="text-gray-700 text-sm md:text-md xl:text-xl 2xl:text-2xl 3xl:text-3xl mb-4 px-2">
          We&apos;ll back soon.
        </p>
        <p className="text-gray-700 text-sm md:text-md xl:text-xl 2xl:text-2xl 3xl:text-3xl mb-4 px-2 font-bold">
          Fernando Agustin Soria, Go Complaint.
        </p>
      </article>
  );
}
