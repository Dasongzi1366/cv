#include "videoio.h"

// VideoWriter
VideoCapture VideoCapture_New() {
    return new cv::VideoCapture();
}

void VideoCapture_Close(VideoCapture v) {
    delete v;
}

bool VideoCapture_Open(VideoCapture v, const char* uri) {
    return v->open(uri);
}

bool VideoCapture_OpenWithAPI(VideoCapture v, const char* uri, int apiPreference) {
    return v->open(uri, apiPreference);
}

bool VideoCapture_OpenWithAPIParams(VideoCapture v, const char* uri, int apiPreference, int *paramsv, int paramsc) {
    std::vector< int > params;

    for( int i = 0; i< paramsc; i++) {
        params.push_back(paramsv[i]);
    }

    return v->open(uri, apiPreference, params);
}

bool VideoCapture_OpenDevice(VideoCapture v, int device) {
    return v->open(device);
}

bool VideoCapture_OpenDeviceWithAPI(VideoCapture v, int device, int apiPreference) {
    return v->open(device, apiPreference);
}

bool VideoCapture_OpenDeviceWithAPIParams(VideoCapture v, int device, int apiPreference, int *paramsv, int paramsc) {
    std::vector< int > params;

    for( int i = 0; i< paramsc; i++) {
        params.push_back(paramsv[i]);
    }

    return v->open(device, apiPreference, params);
}


void VideoCapture_Set(VideoCapture v, int prop, double param) {
    v->set(prop, param);
}

double VideoCapture_Get(VideoCapture v, int prop) {
    return v->get(prop);
}

int VideoCapture_IsOpened(VideoCapture v) {
    return v->isOpened();
}

int VideoCapture_Read(VideoCapture v, Mat buf) {
    return v->read(*buf);
}

void VideoCapture_Grab(VideoCapture v, int skip) {
    for (int i = 0; i < skip; i++) {
        v->grab();
    }
}

int VideoCapture_Retrieve(VideoCapture v, Mat buf) {
    return v->retrieve(*buf);
}

// VideoWriter
VideoWriter VideoWriter_New() {
    return new cv::VideoWriter();
}

void VideoWriter_Close(VideoWriter vw) {
    delete vw;
}

void VideoWriter_Open(VideoWriter vw, const char* name, const char* codec, double fps, int width,
                      int height, bool isColor) {
    int codecCode = cv::VideoWriter::fourcc(codec[0], codec[1], codec[2], codec[3]);
    vw->open(name, codecCode, fps, cv::Size(width, height), isColor);
}

int VideoWriter_IsOpened(VideoWriter vw) {
    return vw->isOpened();
}

void VideoWriter_Write(VideoWriter vw, Mat img) {
    *vw << *img;
}
