# Notes

This package started as a port of [ElPumpo's TinyNvidiaUpdateChecker](https://github.com/ElPumpo/TinyNvidiaUpdateChecker),
but quickly became all encompassing. These are my notes from the porting process.

# TinyNvidiaUpdateChecker 

## Pipline
* Get gpuName and isNotebook by calling the `GetGpuData()` function
* Gets gpuId, osId, and isDchDriver by calling `GetDriverMetadata(gpuName, isNotebook)`
* Calls `GetDriverDownloadInfo(gpuId, osId, isDchDriver)` to oget downloadInfo
* Checks settings for location data and creates download link
* Collects online and offline driver metadata to display to user
* If allowed, downloads newest driver


## Snipets (in C#)

These are code snipets from the base project's pipline. I cut out the garbage to make it more readable.

### Step 1: Get local GPU metadata
```
(string gpuName, bool isNotebook) = GetGpuData();
(int gpuId, int osId, int isDchDriver) = GetDriverMetadata(gpuName, isNotebook);
```

### Step 2: Create download link
```
var downloadInfo = GetDriverDownloadInfo(gpuId, osId, isDchDriver);
var dlPrefix = SettingManager.ReadSetting("Download location");

downloadURL = downloadInfo["DownloadURL"].ToString();
            
// Some GPUs, such as 970M (Win10) URLs (including release notes URL) are HTTP and not HTTPS
if (downloadURL.Contains("https://")) {
    downloadURL = downloadURL.Substring(10);
} else {
    downloadURL = downloadURL.Substring(9);
}

downloadURL = $"https://{dlPrefix}{downloadURL}";
```

### Step 3: Get online driver metadata
```
OnlineGPUVersion = downloadInfo["Version"].ToString();
releaseDate = DateTime.Parse(downloadInfo["ReleaseDateTime"].ToString());
releaseDesc = Uri.UnescapeDataString(downloadInfo["ReleaseNotes"].ToString());

// Cleanup release description
var htmlDocument = new HtmlAgilityPack.HtmlDocument();
htmlDocument.LoadHtml(releaseDesc);

// Remove image nodes
var nodes = htmlDocument.DocumentNode.SelectNodes("//img");
if (nodes != null && nodes.Count > 0)
{
    foreach (var child in nodes) child.Remove();
}

// Remove all links
try {
    var hrefNodes = htmlDocument.DocumentNode.SelectNodes("//a").Where(x => x.Attributes.Contains("href"));
    foreach (var child in hrefNodes) child.Remove();
} catch { }

// Finally set new release description
releaseDesc = htmlDocument.DocumentNode.OuterHtml;

// Get real file size in bytes
using (var request = new HttpRequestMessage(HttpMethod.Head, downloadURL)) {
    using var response = httpClient.Send(request);
    response.EnsureSuccessStatusCode();
    downloadFileSize = response.Content.Headers.ContentLength.Value;
}

// Get PDF release notes
var otherNotes = Uri.UnescapeDataString(downloadInfo["OtherNotes"].ToString());

htmlDocument.LoadHtml(otherNotes);
IEnumerable<HtmlNode> node = htmlDocument.DocumentNode.Descendants("a").Where(x => x.Attributes.Contains("href"));

foreach (var child in node) {
    if (child.Attributes["href"].Value.Contains("release-notes.pdf")) {
        pdfURL = child.Attributes["href"].Value.Trim();
        break;
    }
}

Console.Write("OK!");
Console.WriteLine();

if (debug) {
    Console.WriteLine($"downloadURL: {downloadURL}");
    Console.WriteLine($"pdfURL:      {pdfURL}");
    Console.WriteLine($"releaseDate: {releaseDate.ToShortDateString()}");
    Console.WriteLine($"downloadFileSize:  {Math.Round((downloadFileSize / 1024f) / 1024f)} MiB");
    Console.WriteLine($"OfflineGPUVersion: {OfflineGPUVersion}");
    Console.WriteLine($"OnlineGPUVersion:  {OnlineGPUVersion}");
}
```

### Step 4: Perform driver install if necessary
```
var updateAvailable = false;
var iOffline = int.Parse(OfflineGPUVersion.Replace(".", string.Empty));
var iOnline = int.Parse(OnlineGPUVersion.Replace(".", string.Empty));

if (iOnline == iOffline) {
    Console.WriteLine("There is no new GPU driver available, you are up to date.");
} else if (iOffline > iOnline) {
    Console.WriteLine("Your current GPU driver is newer than what NVIDIA reports!");
} else {
    Console.WriteLine("There is a new GPU driver available to download!");
    updateAvailable = true;
}

if (updateAvailable || forceDL) {
    if (confirmDL) {
        DownloadDriverQuiet(true);
    } else {
        DownloadDriver();
    }
}

Console.WriteLine();
Console.WriteLine("Press any key to exit...");
```