#!/bin/bash
################################################################################
# Help                                                                         #
################################################################################
Help()
{
   # Display Help
   echo "###########################################################"
   echo "Fault: a language and model checker for dynamic systems"
   echo "###########################################################"
   echo
   echo "-f [filepath]     spec file."
   echo "-h                print this help guide."
   echo "-m [mode]         stop compiler at certain milestones: ast,"
   echo "                   ir, smt, or check (default: check)"
   echo
   echo "-c [completeness] check that the system spec is complete"
   echo "                   (default: false)"
   echo 
   echo "-i [input]        format of the input file (default: fspec)"
   echo "-V                print software version and exit."
   echo
}

################################################################################
# Version                                                                        #
################################################################################
Version()
{
   # Lookup the current version of Fault
   docker inspect fault-lang/fault-z3 -f='{{ index .Config.Labels "org.opencontainers.image.version" }}'
}


################################################################################
################################################################################
# Setting up full path (from Docker's perspective)                                                                #
################################################################################
################################################################################

home=${HOME}
pwd=${PWD}
path=${pwd#"$home"}
path=${path#"/"}

################################################################################
################################################################################
# Main program                                                                 #
################################################################################
################################################################################


while getopts f:m:i:c:hV flag
do
    case "${flag}" in
        f) file=${OPTARG};;
        m) mode=${OPTARG};;
        i) input=${OPTARG};;
        c) reach=${OPTARG};;
        h) Help
           exit;;
        V) Version
            exit;;
    esac
done

if [ -z "$file" ]
then
    echo "You must specify a spec file."
else
    # Removing stray '='
    mode=${mode/=/} 
    input=${input/=/}
    reach=${reach/=/}
    file=${file/=/}
    
    filepath="${path}/${file}"

    docker run -v $home:/host:ro fault-lang/fault-z3 -m=$mode -f=$filepath -i=$input -c=$reach
fi
