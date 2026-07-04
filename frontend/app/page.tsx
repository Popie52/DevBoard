"use client";

import { useEffect, useState } from "react";

type HealthResponse = {
  status: string;
  service: string;
  version: string;
};

export default function Home() {
  const [data, setData] = useState<HealthResponse | null>(null);

  useEffect(() => {
    async function load() {
      try {
        const res = await fetch(
          `${process.env.NEXT_PUBLIC_API_URL}/health`
        );
        const json = await res.json();
        setData(json);
      } catch (e) {
        console.error(e);
      }
    }

    load();
  }, []);

  return (
    <main className="min-h-screen relative flex items-center justify-center bg-[#04050A] text-white overflow-hidden">

      {/* chakra aura field */}
      <div className="absolute inset-0">
        <div className="absolute w-[800px] h-[800px] rounded-full bg-orange-500/10 blur-[140px] top-[-300px] left-[-200px]" />
        <div className="absolute w-[700px] h-[700px] rounded-full bg-cyan-400/10 blur-[160px] bottom-[-300px] right-[-250px]" />
        <div className="absolute w-[600px] h-[600px] rounded-full bg-purple-500/10 blur-[180px] top-[20%] left-[35%]" />
      </div>

      {/* subtle energy rings */}
      <div className="absolute w-[900px] h-[900px] border border-white/5 rounded-full animate-pulse" />
      <div className="absolute w-[600px] h-[600px] border border-orange-500/10 rounded-full animate-pulse" />

      <div className="relative w-full max-w-4xl p-6">

        {/* header */}
        <div className="text-center mb-12">
          <h1 className="text-5xl font-light tracking-[0.3em]">
            DevBoard
          </h1>
          <p className="text-white/30 text-sm mt-2">
            chakra system interface
          </p>
        </div>

        {/* bento (solid energy blocks, NOT glass) */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-5">

          {/* CORE STATUS */}
          <div className="md:col-span-2 relative p-8 rounded-2xl
            bg-[#0A0C14] border border-white/10
            shadow-[0_0_80px_-20px_rgba(255,120,0,0.15)]">

            {/* inner chakra glow */}
            <div className="absolute inset-0 flex items-center justify-center">
              <div className="w-[220px] h-[220px] bg-orange-500/10 blur-3xl rounded-full animate-pulse" />
            </div>

            <p className="text-white/40 text-sm relative">
              chakra status
            </p>

            <div className="mt-6 flex items-center gap-4 relative">
              <div className="relative flex h-3 w-3">
                <span className="animate-ping absolute h-full w-full rounded-full bg-orange-400 opacity-60"></span>
                <span className="relative h-3 w-3 rounded-full bg-orange-300"></span>
              </div>

              <span className="text-2xl font-light tracking-wide">
                {data ? data.status : "syncing chakra..."}
              </span>
            </div>

            <p className="mt-6 text-white/25 text-sm relative">
              energy flow from {data?.service ?? "sealed node"}
            </p>
          </div>

          {/* SERVICE */}
          <div className="rounded-2xl p-6 bg-[#0A0C14] border border-white/10
            shadow-[0_0_40px_-20px_rgba(0,200,255,0.12)]">

            <p className="text-white/40 text-sm">node</p>
            <p className="mt-6 text-lg font-light text-cyan-300">
              {data?.service ?? "—"}
            </p>
          </div>

          {/* VERSION */}
          <div className="rounded-2xl p-6 bg-[#0A0C14] border border-white/10
            shadow-[0_0_40px_-20px_rgba(255,0,150,0.10)]">

            <p className="text-white/40 text-sm">chakra flow</p>
            <p className="mt-6 text-lg font-light text-orange-300">
              {data?.version ?? "—"}
            </p>
          </div>

          {/* LORE PANEL */}
          <div className="md:col-span-3 rounded-2xl p-8
            bg-[#0A0C14] border border-white/10
            shadow-[0_0_60px_-25px_rgba(255,120,0,0.10)]">

            <p className="text-white/40 text-sm">chakra network</p>

            <p className="mt-4 text-white/50 text-sm leading-relaxed font-light">
              A sealed system of energy flows beneath the interface.
              Each request releases a controlled pulse through the network,
              maintaining balance across all nodes.
            </p>
          </div>

        </div>
      </div>
    </main>
  );
}