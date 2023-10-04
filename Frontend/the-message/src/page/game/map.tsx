import PlayerSection from "./playerSection";
import Table from "./table";

const Map = () => {
  return (
    <>
      <div className="h-full w-full overflow-hidden" style={{ position: 'relative' }}>
          <Table/>
          <PlayerSection/>
      </div>
    </>
  );
}

export default Map;
