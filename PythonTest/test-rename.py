import os
def rename_file(file_path,file_holder,file_reciever):
    file_full_name=file_path.split("/")[0]
    point1 = file_full_name.split(".")
    double_types=['tar.gz','tar.bz2','tar.bz','tar.Z']
    if len(point1) <= 1:
        filetype = ""
        sp_filename = point1[0]
        new_name = file_holder+"_"+sp_filename+"_"+file_reciever
        #new_name = sp_filename
    elif len(point1) ==2:
        filetype=point1[-1]
        #filename=split_names[0:-1]
        sp_filename = point1[0]
        new_name = file_holder+"_"+sp_filename+"_"+file_reciever+"."+filetype
        #new_name = sp_filename+"."+filetype

    else:
        newtype = point1[-2]+"."+point1[-1]
        if newtype in double_types:
            sp_filename = os.path.splitext(os.path.splitext(file_full_name)[0])[0]
            filetype = newtype
            new_name = file_holder+"_"+sp_filename+"_"+file_reciever+"."+filetype
            #new_name = sp_filename+"."+filetype
        else:
            sp_filename = os.path.splitext(file_full_name)[0]
            filetype = point1[-1]
            new_name = file_holder+"_"+sp_filename+"_"+file_reciever+"."+filetype
            #new_name = sp_filename+"."+filetype
    return new_name
