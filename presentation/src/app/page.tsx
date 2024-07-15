
import MainBanner from "@/components/main/MainBanner";
import MainSection from "@/components/main/MainSection";
import ScrollUpButton from "@/components/main/ScrollUpButton";
import SecondarySection from "@/components/main/SecondarySection";
import TerciarySection from "@/components/main/TerciarySection";

const Home = () => {
  return (
    <div className="w-full relative cursor-default">
      <MainSection />
      <MainBanner />
      <SecondarySection />
      <TerciarySection />
      <ScrollUpButton/>
    </div>
  );
}

export default Home;