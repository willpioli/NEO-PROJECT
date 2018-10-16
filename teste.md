---


---

<p><img src="http://fs1.directupload.net/images/150528/h8n8qwyc.png" alt="CSCore Logo"></p>
<h1 id="csv-data-importer">.csv Data Importer</h1>
<p>Este serviço, desenvolvido em GOLang, tem por objetivo importar dados de um tipo especifico de arquivo .csv<br>
<a href="https://github.com/filoe/cscore.git">https://github.com/filoe/cscore.git</a></p>
<p>You may prefer to install the <a href="https://www.nuget.org/packages/CSCore/">CSCore nuget package</a>:</p>
<pre><code>Install-Package CSCore
</code></pre>
<p>For <strong>FFmpeg</strong> support, install the <a href="https://www.nuget.org/packages/CSCore.Ffmpeg/">CSCore.Ffmpeg nuget package</a></p>
<pre><code>Install-Package CSCore.Ffmpeg -Pre
</code></pre>
<h3 id="why-cscore">Why CSCore?</h3>
<ul>
<li><strong>Highly optimized PERFORMANCE</strong> through usage of CLI instructions</li>
<li><strong>Designed for newbies and professionals</strong></li>
<li><strong>Tons of features</strong></li>
<li><strong>Fast support on <a href="https://github.com/filoe/cscore">github</a>, <a href="http://cscore.codeplex.com/">codeplex</a> or <a href="http://stackoverflow.com/questions/tagged/cscore">stackoverflow</a></strong></li>
<li><strong>High code coverage through unit tests</strong></li>
<li><strong>Licensed under the MS-PL</strong> (does not include the <a href="https://github.com/filoe/cscore/tree/master/CSCore.Ffmpeg">CSCore.Ffmpeg</a> project which is licensed under the LGPL)</li>
</ul>
<h3 id="supported-features">Supported Features</h3>
<p>Currently the following features are implemented:</p>
<ul>
<li><strong>Realtime audio processing</strong>
<ul>
<li>Process audio data in realtime</li>
<li>Apply any processors in any order you want in realtime</li>
<li>Create custom processors (e.g. effects, analyzes, decoders,…)</li>
</ul>
</li>
<li><strong>Codecs</strong> *1
<ul>
<li>MP3</li>
<li>WAVE (PCM, IeeeFloat, GSM, ADPCM,…)</li>
<li>FLAC</li>
<li>AAC</li>
<li>AC3</li>
<li>WMA</li>
<li>Raw data</li>
<li>OGG-Vorbis (through NVorbis)</li>
<li>FFmpeg (lots of additional formats, see <a href="https://github.com/filoe/cscore/tree/master/CSCore.Ffmpeg">CSCore.Ffmpeg</a>)</li>
</ul>
</li>
<li><strong>FFmpeg support</strong>
<ul>
<li>Supported through <a href="https://github.com/filoe/cscore/tree/master/CSCore.Ffmpeg">CSCore.Ffmpeg</a>)</li>
</ul>
</li>
<li><strong>Speaker Output</strong>
<ul>
<li>WaveOut</li>
<li>DirectSoundOut (realtime streaming)</li>
<li>WASAPI (loop- and event-callback + exclusive mode available)</li>
<li>XAudio2</li>
</ul>
</li>
<li><strong>Recording</strong>
<ul>
<li>WaveIn</li>
<li>WASAPI (loop- and event-callback + exclusive mode available)</li>
<li>WASAPI loopback capture (capture output from soundcard)</li>
</ul>
</li>
<li><strong>DSP Algorithms</strong>
<ul>
<li>Fastfouriertransform (FFT)</li>
<li>Effects (Echo, Compressor, Reverb, Chorus, Gargle, Flanger,…)</li>
<li>Resampler</li>
<li>Channel-mixing using custom channel-matrices</li>
<li>Generic Equalizer</li>
<li>…</li>
</ul>
</li>
<li><strong>XAudio2 support</strong>
<ul>
<li>XAudio2.7 and XAudio2.8 support</li>
<li>3D Audio support (see X3DAudio sample)</li>
<li>Streaming source voice implementation allowing<br>
the client to stream any codec, custom effect,… to XAudio2</li>
<li>Optimized for games</li>
</ul>
</li>
<li><strong>Mediafoundation encoding and decoding</strong></li>
<li><strong>DirectX Media Objects Wrapper</strong></li>
<li><strong>CoreAudioAPI Wrapper</strong>
<ul>
<li>WASAPI</li>
<li>Windows Multimedia Devices</li>
<li>Windows Audio Session</li>
<li>Endpoint Volume,…</li>
</ul>
</li>
<li><strong>Multi-Channel support</strong></li>
<li><strong>Flexible</strong>
<ul>
<li>Configure and customize any parts of CSCore</li>
<li>Use low latency values for realtime performance, high latency values for stability</li>
<li>Adjust the audio quality</li>
<li>Configure custom channel matrices</li>
<li>Create custom effects</li>
<li>…</li>
<li>Or simply: <strong>Make CSCore fit your needs!</strong></li>
</ul>
</li>
<li><strong>Tags</strong> (ID3v1, ID3v2, FLAC)</li>
<li><strong>Sample Winforms Visualizations</strong></li>
<li><strong>Optimized performance though the usage of CLI instructions provided by a custom post compiler</strong></li>
</ul>
<p><strong>*1</strong> Some Codecs are only available on certain platforms. For more details, see <a href="http://cscore.codeplex.com/">Codeplex-Page</a>.</p>
<p>Some projects using already using cscore:</p>
<ul>
<li><a href="http://www.digimezzo.com/software/dopamine/">Dopamine</a>: <em>An music player which tries to keep listening to music clean and simple.</em></li>
<li><a href="https://github.com/Alkalinee/Hurricane">Hurricane</a>: <em>Is a powerful music player written in C# based on <a href="https://github.com/filoe/cscore">CSCore sound library</a>.</em></li>
<li><a href="https://github.com/ThuCommix/Sharpex2D">Sharpex2D</a>: A game engine which <em>allows you to create beautiful 2D games under .NET for Windows and Mono compatible systems</em></li>
<li><a href="https://github.com/opcon/turnt-ninja">Turnt-Ninja</a>: A beat-fighting-ninja-like-get-turnt rhythm game inspired by the wonderful Super Hexagon by Terry Cavanagh.</li>
<li><a href="https://www.youtube.com/watch?v=tbrKepBgH3M">HTLED</a>: A audio visualization displayed on a selfmade 32x16 LED matrix.</li>
<li>…</li>
</ul>
<h4 id="samples">Samples:</h4>
<p><a href="Samples/WinformsVisualization">“CSCore - Visualization”</a> Sample:</p>
<p><img src="http://download-codeplex.sec.s-msft.com/Download?ProjectName=cscore&amp;DownloadId=970569" alt="VIS_SAMPLE"></p>
<p><a href="Samples/CSCoreWaveform">“CSCoreWaveform”</a> Sample:</p>
<p><img src="http://fs5.directupload.net/images/160229/adjvd9u9.png" alt="WAVFRM_SAMPLE"></p>
<p>For more samples see <a href="Samples/">Samples</a></p>

